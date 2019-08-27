package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/exception"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/global"
	loginlogic "fgame/fgame/game/login/logic"
	"fgame/fgame/game/login/pbutil"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/welfare/welfare"
	accountypes "fgame/fgame/login/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_LOGIN_TYPE), dispatch.HandlerFunc(handleLogin))
}

//处理登录
func handleLogin(s session.Session, msg interface{}) error {
	log.Debug("处理登陆消息")

	gcs := gamesession.SessionInContext(s.Context())
	//玩家重复登录
	if gcs.State() != gamesession.SessionStateInit {
		log.WithFields(
			log.Fields{
				"playerId": gcs.Player().GetId(),
			}).Warn("login:玩家重复登陆")
		playerlogic.SendSessionSystemMessage(gcs, lang.AccountLoginRepeat)
		return nil
	}

	csLogin := msg.(*uipb.CSLogin)
	token := csLogin.GetToken()
	serverId := csLogin.GetServerId()
	guaJi := csLogin.GetGuaJi()
	deviceMac := csLogin.GetDeviceMac()
	fmt.Println(deviceMac)
	//登陆成功
	err := login(gcs, token, serverId, guaJi)
	if err != nil {
		return err
	}

	return nil
}

func login(gs gamesession.Session, token string, serverId int32, guaJi bool) (err error) {
	originServerId := global.GetGame().GetServerIndex()
	userId, platformUserId, sdkType, deviceType, gm, iosVersion, androidVersion, err := loginlogic.Login(token, originServerId, serverId)
	if err != nil {
		return
	}
	//TODO 验证失败
	if userId == 0 {
		log.WithFields(
			log.Fields{
				"token": token,
			}).Warn("login:用户不存在")
		playerlogic.SendSessionExceptionMessage(gs, exception.ExceptionCodePlayerTokenInvalid)
		gs.Close(true)
		return
	}

	//正常不会出现
	if !sdkType.Valid() {
		log.WithFields(
			log.Fields{
				"token":      token,
				"deviceType": deviceType,
				"sdkType":    sdkType,
			}).Warn("login:sdk无效")
		playerlogic.SendSessionExceptionMessage(gs, exception.ExceptionCodeAccountArgumentInvalid)
		gs.Close(true)
		return
	}

	if !deviceType.Valid() {
		log.WithFields(
			log.Fields{
				"token":      token,
				"deviceType": deviceType,
				"sdkType":    sdkType,
			}).Warn("login:登陆平台错误")
		playerlogic.SendSessionExceptionMessage(gs, exception.ExceptionCodeAccountArgumentInvalid)
		gs.Close(true)
		return
	}

	//TODO:荣昌 验证服务器id

	//获取登陆用户
	p := player.GetOnlinePlayerManager().GetPlayerByUserId(userId)
	if p != nil {
		log.WithFields(
			log.Fields{
				"player": p.GetId(),
			}).Warn("login:用户同时登陆")
		//踢用户
		playerlogic.SendExceptionMessage(p, exception.ExceptionCodePlayerLoginSameTime)
		p.Close(nil)
		<-p.Done()
	}

	state := accountypes.RealNameStateUp18
	now := global.GetGame().GetTimeService().Now()
	openTime := global.GetGame().GetServerTime()
	mergeTime := merge.GetMergeService().GetMergeTime()
	activityOpenTime := welfare.GetWelfareService().GetServerStartTime()  //global.GetGame().GetServerTime()
	activityMergeTime := welfare.GetWelfareService().GetServerMergeTime() //merge.GetMergeService().GetMergeTime()
	if openTime > now && gm == 0 {
		log.WithFields(
			log.Fields{
				"userId": userId,
			}).Warn("login:系统还未开始登陆")
		openTimeTime := timeutils.MillisecondToTime(openTime)
		openTimeStr := openTimeTime.Format("2006-01-02 15:04:05")
		exceptionContent := fmt.Sprintf(lang.GetLangService().ReadLang(exception.ExceptionCodeServerNoOpen.LangCode()), openTimeStr)
		playerlogic.SendSessionExceptionContentMessage(gs, exceptionContent)
		gs.Close(true)
		return
	}

	scLogin := pbutil.BuildSCLogin(userId, accountypes.RealNameStateUp18, now, openTime, mergeTime, activityOpenTime, activityMergeTime, gm, iosVersion, androidVersion)
	gs.Send(scLogin)
	//加载玩家数据
	err = playerlogic.LoadPlayer(gs, sdkType, deviceType, platformUserId, serverId, userId, state, guaJi, gm != 0)
	if err != nil {
		return
	}

	return
}
