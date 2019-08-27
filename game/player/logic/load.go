package logic

import (
	"context"
	logintypes "fgame/fgame/account/login/types"
	"fgame/fgame/common/exception"
	"fgame/fgame/common/lang"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/dao"
	playerentity "fgame/fgame/game/player/entity"
	"fgame/fgame/game/player/operation"
	"fgame/fgame/game/player/pbutil"
	playerplayer "fgame/fgame/game/player/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/register/register"
	"fgame/fgame/game/scene/scene"
	gamesession "fgame/fgame/game/session"

	gameevent "fgame/fgame/game/event"
	playereventtypes "fgame/fgame/game/player/event/types"
	accounttypes "fgame/fgame/login/types"
	"fgame/fgame/pkg/idutil"

	log "github.com/Sirupsen/logrus"
)

//根据平台id加载用户
func LoadPlayer(gs gamesession.Session, sdkType logintypes.SDKType, deviceType logintypes.DevicePlatformType, platformUserId string, serverId int32, userId int64, realNameState accounttypes.RealNameState, guaJi bool, gm bool) (err error) {
	p := playerplayer.NewPlayer(gs, sdkType, deviceType, platformUserId, serverId, userId, realNameState, guaJi, gm)
	//session认证
	flag := gs.Auth(p)
	if !flag {
		log.WithFields(
			log.Fields{
				"userId": userId,
			}).Error("player:用户认证超时")
		SendSessionExceptionMessage(gs, exception.ExceptionCodePlayerLoginTimeout)
		gs.Close(true)
		return
	}

	//添加玩家失败
	flag = player.GetOnlinePlayerManager().AddPlayer(p)
	if !flag {
		log.WithFields(
			log.Fields{
				"userId": userId,
			}).Error("player:用户进入服务器,失败")
		//踢用户
		SendSessionSystemMessage(gs, lang.AccountLoginFailed)
		gs.Close(true)
		return
	}

	log.WithFields(
		log.Fields{
			"userId": userId,
		}).Info("player:用户进入服务器")

	flag = p.EnterAuth()
	if !flag {
		//状态改变错误
		log.WithFields(
			log.Fields{
				"userId": userId,
			}).Error("player:用户进入认证,失败")
		SendSessionExceptionMessage(gs, exception.ExceptionCodePlayerStateException)
		gs.Close(true)
		return
	}

	log.WithFields(
		log.Fields{
			"userId": userId,
		}).Info("player:用户进入认证成功")
	flag = p.EnterLoadingRoleList()
	if !flag {
		//状态改变错误
		log.WithFields(
			log.Fields{
				"userId": userId,
			}).Error("player:用户加载角色列表,失败")
		SendSessionExceptionMessage(gs, exception.ExceptionCodePlayerStateException)
		gs.Close(true)
		return
	}

	pe, err := dao.GetPlayerDao().QueryByUserId(userId, serverId)
	if err != nil {
		return
	}

	if pe == nil {
		if !register.GetRegisterService().IsOpenRegister() {
			log.WithFields(
				log.Fields{
					"userId": userId,
				}).Warn("player:处理创建角色消息,关闭注册")
			SendSessionExceptionMessage(gs, exception.ExceptionCodeRegisterClose)
			gs.Close(true)
			return
		}
		//等候选择角色
		flag = p.EnterWaitingSelectRole()
		if !flag {
			//状态改变错误
			log.WithFields(
				log.Fields{
					"userId": userId,
				}).Error("player:用户进入等候选择角色,失败")
			SendSessionExceptionMessage(gs, exception.ExceptionCodePlayerStateException)
			gs.Close(true)
			return
		}
		log.WithFields(
			log.Fields{
				"userId": userId,
			}).Info("player:用户进入选择角色")
		scEnterSelectJob := pbutil.BuildSCEnterSelectJob()
		//推送进入选择角色
		gs.Send(scEnterSelectJob)
		return
	}
	log.WithFields(
		log.Fields{
			"userId": userId,
		}).Info("player:用户加载角色列表成功")
	now := global.GetGame().GetTimeService().Now()
	//禁号判断
	if pe.Forbid == 1 {
		if pe.ForbidEndTime == 0 || (pe.ForbidEndTime > now) {
			log.WithFields(
				log.Fields{
					"userId": userId,
				}).Warn("player:玩家禁号")
			//踢用户
			SendSessionExceptionContentMessage(gs, pe.ForbidText)
			gs.Close(true)
			return
		}
	}

	//触发选择角色
	flag = p.Auth(pe.Id)
	if !flag {
		log.WithFields(
			log.Fields{
				"userId": userId,
			}).Error("player:用户选择角色,失败")
		SendSessionExceptionMessage(gs, exception.ExceptionCodePlayerStateException)
		gs.Close(true)
		return
	}
	log.WithFields(
		log.Fields{
			"userId": userId,
		}).Info("player:用户进入角色成功")

	//加载角色
	loadAll(p)

	return
}

//TODO 合服后创建角色问题
//创建角色
func CreatePlayer(p player.Player, role playertypes.RoleType, sex playertypes.SexType, name string) (err error) {
	log.WithFields(
		log.Fields{
			"userId": p.GetUserId(),
		}).Info("player:用户正在创建角色")
	userId := p.GetUserId()
	//当前服务器
	serverIdIndex := global.GetGame().GetServerIndex()
	//原始服务器
	serverId := p.GetServerId()
	//sdk类型
	sdkType := p.GetSDKType()

	pe, err := dao.GetPlayerDao().QueryByUserId(userId, serverId)
	if err != nil {
		return
	}
	if pe != nil {
		log.WithFields(
			log.Fields{
				"userId": p.GetUserId(),
			}).Warn("player:用户重复创建角色")
		SendSystemMessage(p, lang.RepeatCreateJob)
		return
	}
	pe, err = dao.GetPlayerDao().QueryByName(serverId, name)
	if err != nil {
		return
	}
	if pe != nil {
		log.WithFields(
			log.Fields{
				"userId": p.GetUserId(),
				"name":   name,
			}).Warn("player:名字已经存在")
		SendSystemMessage(p, lang.NameAlreadyExist)
		return
	}
	flag := p.EnterCreateRole()
	if !flag {
		log.WithFields(
			log.Fields{
				"userId": p.GetUserId(),
			}).Error("player:用户进入创建角色,失败")
		SendExceptionMessage(p, exception.ExceptionCodePlayerStateException)
		p.Close(nil)
		return
	}

	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	defaultSystemCompensate := int32(1)
	pe = &playerentity.PlayerEntity{
		Id:               id,
		UserId:           userId,
		ServerId:         serverIdIndex,
		SdkType:          int32(sdkType),
		OriginServerId:   serverId,
		Name:             name,
		Role:             int32(role),
		Sex:              int32(sex),
		CreateTime:       now,
		SystemCompensate: defaultSystemCompensate,
	}
	if err = global.GetGame().GetDB().DB().Save(pe).Error; err != nil {
		return
	}
	log.WithFields(
		log.Fields{
			"userId": p.GetUserId(),
		}).Info("player:用户创建角色成功")
	//返回选择角色成功
	scSelectJob := pbutil.BuildSCSelectJob()
	p.SendMsg(scSelectJob)

	flag = p.Auth(pe.Id)
	if !flag {
		//状态改变错误
		log.WithFields(
			log.Fields{
				"userId": userId,
			}).Error("player:用户进入角色,失败")
		SendExceptionMessage(p, exception.ExceptionCodePlayerStateException)
		p.Close(nil)
		return
	}

	log.WithFields(
		log.Fields{
			"userId": p.GetUserId(),
		}).Info("player:用户进入角色")
	//加载角色
	loadAll(p)

	return
}

//异步加载
func loadAll(p player.Player) {
	flag := p.EnterLoad()
	if !flag {
		log.WithFields(
			log.Fields{
				"userId": p.GetUserId(),
			}).Error("player:玩家进入加载角色,失败")
		SendExceptionMessage(p, exception.ExceptionCodePlayerStateException)
		p.Close(nil)
		return
	}

	//发布异步加载操作
	os := global.GetGame().GetOperationService()
	//TODO 荣昌 是否需要修改context
	o := operation.CreateLoadPlayerOperation(p.Session().Context(), onPlayerLoadAll)
	os.PostOperation(o)
	return
}

//回调
func onPlayerLoadAll(ctx context.Context, result interface{}, err error) error {
	//TODO 捕捉异常

	s := gamesession.SessionInContext(ctx)
	p := s.Player().(player.Player)

	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
				"err":      err,
			}).Error("player:用户加载用户数据,错误")
		SendExceptionMessage(p, exception.ExceptionCodePlayerStateException)
		p.Close(nil)
		return nil
	}

	// flag := result.(bool)
	// if !flag {
	// 	log.WithFields(
	// 		log.Fields{
	// 			"playerId": p.GetId(),
	// 		}).Error("player:用户加载用户数据,失败")
	// 	SendExceptionMessage(p, exception.ExceptionCodePlayerStateException)
	// 	p.Close(nil)
	// 	return nil
	// }

	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
		}).Info("player:用户加载用户数据,完成")
	gameevent.Emit(playereventtypes.EventTypePlayerLoadFinish, p, nil)
	return nil
}

func SendSystemMessage(pl scene.Player, code lang.LangCode, args ...string) {
	content := lang.GetLangService().ReadLang(code)
	scSystemMessage := pbutil.BuildSCSystemMessage(content, args...)
	pl.SendMsg(scSystemMessage)

	return
}

func SendSystemContentMessage(pl scene.Player, content string) {
	scSystemMessage := pbutil.BuildSCSystemMessage(content)
	pl.SendMsg(scSystemMessage)
	return
}

func SendSessionSystemMessage(gs gamesession.Session, code lang.LangCode, args ...string) {
	content := lang.GetLangService().ReadLang(code)
	scSystemMessage := pbutil.BuildSCSystemMessage(content, args...)
	gs.Send(scSystemMessage)
	return
}

func SendExceptionMessage(pl player.Player, code exception.ExceptionCode) {
	content := lang.GetLangService().ReadLang(code.LangCode())
	scException := pbutil.BuildSCException(content, code)
	pl.SendMsg(scException)

	return
}

func SendExceptionContentMessage(pl player.Player, content string) {
	code := exception.ExceptionCodeKickout
	scException := pbutil.BuildSCException(content, code)
	pl.SendMsg(scException)
	return
}

func SendSessionExceptionMessage(gs gamesession.Session, code exception.ExceptionCode) {
	content := lang.GetLangService().ReadLang(code.LangCode())
	scException := pbutil.BuildSCException(content, code)
	gs.Send(scException)
	return
}

func SendSessionExceptionContentMessage(gs gamesession.Session, content string) {
	code := exception.ExceptionCodeKickout
	scException := pbutil.BuildSCException(content, code)
	gs.Send(scException)
	return
}
