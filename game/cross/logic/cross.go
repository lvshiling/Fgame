package logic

import (
	"context"
	"errors"
	centertypes "fgame/fgame/center/types"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	"fgame/fgame/core/session"
	grpcsession "fgame/fgame/core/session/grpc"
	grpcpb "fgame/fgame/core/session/grpc/pb"
	"fgame/fgame/game/center/center"
	gamecodec "fgame/fgame/game/codec"
	"fgame/fgame/game/cross/cross"
	crosseventtypes "fgame/fgame/game/cross/event/types"
	"fgame/fgame/game/cross/pbutil"
	playercross "fgame/fgame/game/cross/player"
	crosssession "fgame/fgame/game/cross/session"
	crosstypes "fgame/fgame/game/cross/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	loginlogic "fgame/fgame/game/login/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//用户进入跨服
func PlayerEnterCross(pl player.Player, crossType crosstypes.CrossType, args ...string) {
	PlayerEnterCrossWithBehavior(pl, crossType, crosstypes.CrossBehaviorTypeNormal, args...)
	return
}

func PlayerEnterCrossWithBehavior(pl player.Player, crossType crosstypes.CrossType, behaviorType crosstypes.CrossBehaviorType, args ...string) {

	crossManager := pl.GetPlayerDataManager(playertypes.PlayerCrossDataManagerType).(*playercross.PlayerCrossDataManager)
	tempArgs := make([]string, 0, len(args)+1)
	tempArgs = append(tempArgs, fmt.Sprintf("%d", behaviorType))
	tempArgs = append(tempArgs, args...)
	crossManager.EnterCross(crossType, tempArgs...)

	//异步连接跨服
	connectCross(pl, crossType.GetServerType())
	return
}

func PlayerTracEnterCross(pl player.Player, crossType crosstypes.CrossType, behaviorType crosstypes.CrossBehaviorType, foeId string) (flag bool) {
	if !CheckBeforeEnterCross(pl, crossType) {
		return
	}

	PlayerEnterCrossWithBehavior(pl, crossType, behaviorType, foeId)
	return
}

func CheckBeforeEnterCross(pl player.Player, crossType crosstypes.CrossType) bool {
	if !cross.CheckEnterCross(pl, crossType) {
		return false
	}

	return true
}

func PlayerReenterCross(pl player.Player, crossType crosstypes.CrossType) {
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Info("cross:用户重新进入跨服")
	//跨服不重连
	switch crossType {
	case crosstypes.CrossTypeLianYu,
		crosstypes.CrossTypeGodSiegeQiLin,
		crosstypes.CrossTypeGodSiegeHuoFeng,
		crosstypes.CrossTypeGodSiegeDuLong,
		crosstypes.CrossTypeDenseWat,
		crosstypes.CrossTypeShenMoWar:
		PlayerExitCross(pl)
		return
	}
	//异步连接跨服
	connectCross(pl, crossType.GetServerType())
	return
}

//用户退出跨服和返回上一个场景
func PlayerExitCross(pl player.Player) {
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Info("cross:跨服数据退出")
	crossManager := pl.GetPlayerDataManager(playertypes.PlayerCrossDataManagerType).(*playercross.PlayerCrossDataManager)
	crossManager.ExitCross()
	crossSession := pl.GetCrossSession()
	if crossSession != nil {
		crossSession.Close(true)
	}
}

//用户退出跨服和返回上一个场景
func AsyncPlayerExitCross(pl player.Player) {
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Info("arena:跨服数据异步退出")
	ctx := scene.WithPlayer(context.Background(), pl)

	pl.Post(message.NewScheduleMessage(onAsyncPlayerExitCross, ctx, nil, nil))
}

func onAsyncPlayerExitCross(ctx context.Context, result interface{}, err error) error {
	pl := scene.PlayerInContext(ctx)
	p := pl.(player.Player)
	PlayerExitCross(p)
	return nil
}

var (
	errorCrossConnNoFound = errors.New("跨服连接不存在")
	errorCrossNoSupport   = errors.New("跨服不支持")
)

//连接跨服
func connectCross(pl player.Player, serverType centertypes.GameServerType) {
	crossDisable := global.GetGame().CrossDisable()
	if crossDisable {
		err := errorCrossNoSupport
		ctx := scene.WithPlayer(context.Background(), pl)
		//回调连接失败
		sm := message.NewScheduleMessage(onPlayerCrossConnect, ctx, nil, err)
		pl.Post(sm)
		return
	}

	go func() {
		var err error
		defer func() {
			//TODO 捕捉panic
			if err != nil {
				ctx := scene.WithPlayer(context.Background(), pl)
				//回调连接失败
				sm := message.NewScheduleMessage(onPlayerCrossConnect, ctx, nil, err)
				//跨服携程处理
				// cross.GetCrossService().GetCross().Post(sm)
				pl.Post(sm)
			}
		}()
		//获取grpc连接
		conn := center.GetCenterService().GetCross(serverType)
		if conn == nil {
			err = errorCrossConnNoFound
			return
		}

		openHandler := session.SessionHandlerFunc(onSessionOpen)
		closeHandler := session.SessionHandlerFunc(onSessionClose)
		receiveHandler := session.HandlerFunc(onSessionReceive)
		sendHandler := session.HandlerFunc(onSessionSend)

		h := grpcsession.NewGrpcClientHandler(openHandler, closeHandler, receiveHandler, sendHandler)
		cc := grpcpb.NewConnectionClient(conn)
		//TODO 制作超时和元信息
		ctx := pl.GetContext()
		ccc, err := cc.Connect(ctx)
		if err != nil {
			//TODO 连接失败
			return
		}
		h.Handle(ccc)
	}()
}

//跨服成功回调
func onPlayerCrossConnect(ctx context.Context, result interface{}, err error) (rerr error) {
	pl := scene.PlayerInContext(ctx)
	p := pl.(player.Player)
	if err != nil {
		onPlayerCrossConnectFailed(p, err)
		return nil
	}
	sess := result.(session.Session)
	onPlayerCrossConnectSuccess(p, sess)

	return
}

//跨服连接失败
func onPlayerCrossConnectFailed(pl player.Player, err error) {
	crossType := pl.GetCrossType()
	//不管奔溃都要移除
	defer OnPlayerExitCross(pl, crossType)

	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
			}).Warn("cross:玩家跨服,失败")
	}
	//TODO 重试
	playerlogic.SendSystemMessage(pl, lang.CrossFailed)

	//退出
	crossManager := pl.GetPlayerDataManager(playertypes.PlayerCrossDataManagerType).(*playercross.PlayerCrossDataManager)
	crossManager.ExitCross()

	return
}

//跨服连接成功
func onPlayerCrossConnectSuccess(pl player.Player, sess session.Session) {
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Info("cross:玩家跨服,连接成功")
	//创建跨服对话和上下文
	ps := crosssession.NewSendSession(sess, gamecodec.GetCodec(), nil)
	nctx := crosssession.WithSendSession(pl.GetContext(), ps)
	sess.SetContext(nctx)
	//TODO 跨服管理器添加用户
	//设置跨服session
	pl.SetCrossSession(ps)

	//TODO 发送连接成功事件
	//发送登陆消息
	siLogin := pbutil.BuildSILogin(pl.GetId())
	pl.SendCrossMsg(siLogin)
}

//跨服连接关闭
// func onPlayerCrossConnectClose(ctx context.Context, result interface{}, err error) (rerr error) {
// 	pl := scene.PlayerInContext(ctx)
// 	p := pl.(player.Player)
// 	log.WithFields(
// 		log.Fields{
// 			"playerId": pl.GetId(),
// 		}).Info("cross:跨服连接关闭")
// 	//TODO 被动断开 重连

// 	//退出跨服
// 	onPlayerExitCross(p)

// 	return
// }

var (
	crossFailed = errors.New("cross failed")
)

//对话打开
func onSessionOpen(sess session.Session) (err error) {
	//设置到玩家身上
	gameS := gamesession.SessionInContext(sess.Context())
	pl := gameS.Player().(player.Player)
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Info("cross:跨服,对话打开")
	ctx := scene.WithPlayer(context.Background(), pl)
	sm := message.NewScheduleMessage(onPlayerCrossConnect, ctx, sess, nil)
	pl.Post(sm)
	return nil
}

//对话关闭
func onSessionClose(sess session.Session) (err error) {

	gameS := gamesession.SessionInContext(sess.Context())
	pl := gameS.Player().(player.Player)
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Info("cross:跨服,对话关闭")
	pl.LogoutCross()
	// ctx := scene.WithPlayer(context.Background(), pl)

	// sm := message.NewScheduleMessage(onPlayerCrossConnectClose, ctx, nil, nil)
	// pl.Post(sm)
	// cross.GetCrossService().GetCross().Post(sm)
	return nil
}

//接受消息
func onSessionReceive(sess session.Session, msg []byte) (err error) {
	return processor.GetMessageProcessor().ProcessCross(sess, msg)
}

//发送消息
func onSessionSend(sess session.Session, msg []byte) (err error) {
	return nil
}

//跨服关闭
func OnPlayerExitCross(pl player.Player, crossType crosstypes.CrossType) {
	//退出匹配状态
	gameevent.Emit(crosseventtypes.EventTypePlayerCrossExit, pl, crossType)
	//移除跨服
	cross.GetCrossService().GetCross().RemovePlayer(pl)
	if pl.IsCross() {
		//防止退出奔溃
		defer func() {
			//捕捉奔溃
			if r := recover(); r != nil {
				log.WithFields(
					log.Fields{
						"playerId": pl.GetId(),
						"error":    r,
					}).Error("cross:退出跨服,异常")
				//登出中
				if pl.IsLogouting() {
					loginlogic.Logout(pl)
				} else {
					pl.Close(nil)
				}
				return
			}

			//登出
			if pl.IsLogouting() {
				loginlogic.Logout(pl)
			}
		}()

		flag := pl.LeaveCross()
		if !flag {
			panic(fmt.Errorf("cross:退出跨服应该成功"))
		}

		if !pl.IsLogouting() {
			scenelogic.PlayerEnterOriginScene(pl)
		}
	} else {
		//登出
		if pl.IsLogouting() {
			//退出场景
			loginlogic.Logout(pl)
			return
		}
	}

	return
}

//玩家数据推送
func CrossPlayerDataLogin(pl player.Player) {
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Info("cross:跨服,跨服数据推送")

	//退出场景
	if pl.GetScene() != nil {
		//退出场景 失败
		scenelogic.PlayerExitScene(pl, true)
	}

	flag := pl.EnterCross()
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("cross:跨服,进入跨服失败")
		se := pl.GetCrossSession()
		if se != nil {
			se.Close(true)
		}
		return
	}
	ctx := scene.WithPlayer(context.Background(), pl)
	sm := message.NewScheduleMessage(onPlayerEnterCross, ctx, nil, nil)
	cross.GetCrossService().GetCross().Post(sm)
}

func onPlayerEnterCross(ctx context.Context, result interface{}, err error) (rerr error) {
	p := scene.PlayerInContext(ctx)
	pl := p.(player.Player)
	cross.GetCrossService().GetCross().AddPlayer(pl)
	flag := pl.Cross()
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("cross:跨服,进入跨服登陆失败")
		//TODO 发送异常代码
		pl.GetCrossSession().Close(true)
		return
	}

	gameevent.Emit(crosseventtypes.EventTypePlayerCrossEnter, pl, nil)
	//推送用户数据
	siPlayerData := pbutil.BuildSIPlayerData(pl)
	//推送用户数据
	pl.SendCrossMsg(siPlayerData)
	return
}
