package monitor

import (
	"context"
	"fgame/fgame/gm/gamegm/basic/pb"
	"fgame/fgame/gm/gamegm/pkg/timeutils"
	"fgame/fgame/gm/gamegm/session"
	"strconv"

	login "fgame/fgame/gm/gamegm/gm/user/service"

	loginpb "fgame/fgame/gm/gamegm/monitor/login/pb"
	messagetypepb "fgame/fgame/gm/gamegm/monitor/messagetype/pb"
	"fgame/fgame/gm/gamegm/monitor/model"

	chatpb "fgame/fgame/gm/gamegm/monitor/chatmonitor/pb/chat"
	chatmessagetypepb "fgame/fgame/gm/gamegm/monitor/chatmonitor/pb/messagetype"
	logserverlog "fgame/fgame/logserver/log"
	logservermodel "fgame/fgame/logserver/model"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

const (
	monitorServiceKey = contextKey("monitor.service")
)

func MonitorServiceInContext(ctx context.Context) *MonitorService {
	s, ok := ctx.Value(monitorServiceKey).(*MonitorService)
	if !ok {
		return nil
	}
	return s
}

func WithMonitorService(ctx context.Context, s *MonitorService) context.Context {
	return context.WithValue(ctx, monitorServiceKey, s)
}

type MonitorService struct {
	stats      *stats
	pm         *PlayerManager
	dis        *Dispatcher
	los        login.ILoginService
	userManage IUserServerManager
	_done      chan struct{}
	cen        ICenterServer
}

func NewMonitorService(p_dis *Dispatcher, p_los login.ILoginService, p_center ICenterServer) *MonitorService {
	rst := &MonitorService{
		dis: p_dis,
		los: p_los,
	}
	rst.stats = newStats()
	rst.pm = NewPlayerManager()
	rst.userManage = NewUserServerManage()
	rst._done = make(chan struct{})
	rst.cen = p_center
	return rst
}

//对话开启
func (qps *MonitorService) SessionOpen(s session.Session) error {
	//设置玩家
	// uo := &UserOptions{
	// 	AuthTimeout:  qps.opts.AuthTimeout,
	// 	PingTimeout:  qps.opts.PingTimeout,
	// 	SendTimeout:  qps.opts.SendTimeout,
	// 	SendMsgQueue: qps.opts.SendMsgQueue,
	// }

	log.WithFields(
		log.Fields{
			"sessionId": s.Id(),
		}).Debug("SessionOpen,接收消息")
	u := NewUser(s, nil)
	nctx := WithMonitorService(s.Context(), qps)
	nctx = WithPlayer(nctx, u)
	s.SetContext(nctx)
	handleSessionOpenStats(s)
	return nil
}

//对话关闭
func (qps *MonitorService) SessionClose(s session.Session) error {

	handleSessionCloseStats(s)

	pl := PlayerInContext(s.Context())
	if pl == nil {
		panic("SessionClose: never reach here")
	}

	pl.Close()
	qps.pm.RemovePlayer(pl)
	qps.userManage.ClearUser(pl.Id())
	return nil
}

func (qps *MonitorService) SessionReceive(s session.Session, msg []byte) error {

	handleSessionRecvStats(s, msg)
	log.WithFields(
		log.Fields{
			"sessionId": s.Id(),
			"msg":       msg,
		}).Debug("对话处理器,接收消息")
	m := &pb.Message{}
	err := proto.Unmarshal(msg, m)
	if err != nil {
		return err
	}

	//TODO 怎么处理消息 根据用户状态

	// pl := PlayerInContext(s.Context())
	if !isWorkMessage(m.GetMessageType()) { //非工作线程的消息，如果有工作县城的消息这边需要建立工作线程处理
		return qps.dis.Handle(s, m)
	}
	return nil
}

func (qps *MonitorService) SessionSend(s session.Session, msg []byte) error {
	return handleSessionSendStats(s, msg)
}

func (qps *MonitorService) Login(pl Player, token string) (playerId int64, err error) {

	playerId, _, err = qps.los.VerifyToken(token)
	//验证token失败
	if err != nil {
		return
	}

	tpl := qps.pm.GetPlayerById(playerId)
	//在服���器内
	if tpl != nil {
		err = ErrorPlayerLoginSameTime
		CloseWithError(tpl.Session(), err)
		return
	}

	//创建玩家
	flag := pl.Auth(playerId, pl.Session().Ip())
	if !flag {
		err = ErrorPlayerAuthTimeout
		return
	}
	//添加用户
	err = qps.pm.AddPlayer(pl)
	if err != nil {
		return
	}

	//发送登录消息
	gcLogin := &loginpb.GCLogin{}
	gcLogin.PlayerId = &playerId
	gcMsg := &pb.Message{}

	gcMsgType := int32(messagetypepb.QiPaiMessageType_GCLoginType)
	gcMsg.MessageType = &gcMsgType
	err = proto.SetExtension(gcMsg, loginpb.E_GcLogin, gcLogin)
	if err != nil {
		return
	}

	gcMsgB, err := proto.Marshal(gcMsg)
	if err != nil {
		return
	}
	pl.Send(gcMsgB)
	return
}

//日志监测服务接口
func (qps *MonitorService) HandleLog(msg logserverlog.LogMsg) (err error) {
	chatmst, ok := msg.(*logservermodel.ChatContent)
	if !ok || chatmst == nil {
		log.Debug("聊天服务中：强制转换异常")
		return nil
	}

	log.WithFields(log.Fields{
		"server":           chatmst.ServerId,
		"platform":         chatmst.Platform,
		"chatmst.Channel":  chatmst.Channel,
		"chatmst.Ip":       chatmst.Ip,
		"chatmst.RecvId":   chatmst.RecvId,
		"chatmst.RecvName": chatmst.RecvName,
		"servertype":       chatmst.ServerType,
	}).Debug("收到的日志")
	if chatmst.ServerType != 0 {
		return nil
	}

	chatInfo := &model.ChatMsg{
		PlayerId:         chatmst.PlayerId,
		PlayerName:       chatmst.Name,
		VipLevel:         chatmst.Vip,
		GameLevel:        chatmst.Level,
		ChatType:         chatmst.Channel,
		ChatMethod:       chatmst.MsgType,
		ChatMsg:          string(chatmst.Content),
		ChatTime:         chatmst.LogTime,
		ToPlayerId:       chatmst.RecvId,
		ToPlayerName:     chatmst.RecvName,
		Ip:               chatmst.Ip,
		CenterPlatformId: chatmst.Platform,
		CenterServerId:   chatmst.ServerId,
		UserId:           chatmst.UserId,
	}

	serverid := qps.cen.GetCenterServerDBId(chatmst.Platform, chatmst.ServerId)
	if serverid < 1 {
		return
	}
	qps.send(int32(serverid), chatInfo)
	return nil
}

func (qps *MonitorService) send(p_serverId int32, chatInfo *model.ChatMsg) error {
	log.Debug("发送玩家谈话")
	userList := qps.userManage.GetServerUserList(p_serverId)
	log.Debug("服务器里的玩家数：", len(userList))
	if userList == nil || len(userList) == 0 {
		return nil
	}

	for _, value := range userList {
		play := qps.pm.GetPlayerById(value)
		log.Debug("玩家ID:", value)
		if play == nil {
			log.Debug("获取玩家为空")
			continue
		}

		playerId := strconv.FormatInt(chatInfo.PlayerId, 10)
		userId := strconv.FormatInt(chatInfo.UserId, 10)
		// playerId := chatInfo.PlayerId
		toPlayerId := strconv.FormatInt(chatInfo.ToPlayerId, 10)
		gcUserServer := &chatpb.GCChatMinitorMsg{}
		gcUserServer.PlayerId = &playerId
		gcUserServer.PlayerName = &chatInfo.PlayerName
		gcUserServer.ToPlayerId = &toPlayerId
		gcUserServer.ToPlayerName = &chatInfo.ToPlayerName
		gcUserServer.VipLevel = &chatInfo.VipLevel
		gcUserServer.ChatMethod = &chatInfo.ChatMethod
		gcUserServer.ChatMsg = &chatInfo.ChatMsg
		gcUserServer.ChatTime = &chatInfo.ChatTime
		gcUserServer.ChatType = &chatInfo.ChatType
		gcUserServer.GameLevel = &chatInfo.GameLevel
		gcUserServer.CenterPlatformId = &chatInfo.CenterPlatformId
		gcUserServer.CenterServerId = &chatInfo.CenterServerId
		gcUserServer.Ip = &chatInfo.Ip
		gcUserServer.UserId = &userId
		gcMsg := &pb.Message{}
		gcMsgType := int32(chatmessagetypepb.ChatMonitorMessageType_GCChatMinitorMsgType)
		gcMsg.MessageType = &gcMsgType
		err := proto.SetExtension(gcMsg, chatpb.E_GcChatMinitorMsg, gcUserServer)
		if err != nil {
			log.Debug("发送谈话，proto设置拓展异常", err)
			CloseWithError(play.Session(), err)
			continue
		}
		gcMsgB, err := proto.Marshal(gcMsg)
		if err != nil {
			log.Debug("发送谈话，proto序列化异常")
			continue
		}
		play.Send(gcMsgB)
	}
	return nil
}

func (qps *MonitorService) StartTestTick() {
	go func() {
	loop:
		for {
			select {
			case <-time.After(time.Minute * 1):
				{
					qps.cen.SyncServer()
				}
			case <-qps._done:
				{
					break loop
				}
			}
		}
	}()
}

func (qps *MonitorService) testSendChat() {
	qps.userManage.Log()
	info := &model.ChatMsg{
		PlayerId:     1000000000,
		PlayerName:   "小毛驴",
		VipLevel:     99,
		GameLevel:    180,
		ChatType:     1,
		ChatMethod:   1,
		ChatMsg:      "我有一头小毛驴呀从来都不骑，啦啦啦....",
		ChatTime:     timeutils.TimeToMillisecond(time.Now()),
		ToPlayerId:   200,
		ToPlayerName: "大毛驴",
	}
	qps.send(1, info)
}

const (
	minWookMessageType = 2000
	maxWookMessageType = 3000
)

func (qps *MonitorService) GetUserServerManage() IUserServerManager {
	return qps.userManage
}

//是否工作信息
func isWorkMessage(msgType int32) bool {
	if msgType < minWookMessageType {
		return false
	}
	if msgType > maxWookMessageType {
		return false
	}
	return true
}
