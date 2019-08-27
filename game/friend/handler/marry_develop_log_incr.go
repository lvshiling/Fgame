package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/friend/friend"
	"fgame/fgame/game/friend/pbutil"
	playerfriend "fgame/fgame/game/friend/player"
	friendtypes "fgame/fgame/game/friend/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FRIEND_MARRY_DEVELOP_LOG_INCR_TYPE), dispatch.HandlerFunc(handleMarryDevelopLogIncr))
}

//处理表白日志请求
func handleMarryDevelopLogIncr(s session.Session, msg interface{}) (err error) {
	log.Debug("friend:处理获取表白日志请求")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csLogIncr := msg.(*uipb.CSFriendMarryDevelopLogIncr)
	logTime := csLogIncr.GetLogTime()
	logTypeInt := csLogIncr.GetLogType()

	//参数不对
	logType := friendtypes.MarryDevelopLogType(logTypeInt)
	if !logType.Valid() {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"logTypeInt": logTypeInt,
			}).Warn("friend:处理获取表白日志类型,错误")
		return
	}

	err = marryDevelopLogIncr(tpl, logType, logTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"logTypeInt": logTypeInt,
				"logTime":    logTime,
				"error":      err,
			}).Error("friend:处理获取表白日志请求,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"logTypeInt": logTypeInt,
			"logTime":    logTime,
		}).Debug("friend:处理获取表白日志请求完成")
	return nil

}

//获取表白记录界面信息逻辑
func marryDevelopLogIncr(pl player.Player, logType friendtypes.MarryDevelopLogType, logTime int64) (err error) {
	friendManager := pl.GetPlayerDataManager(playertypes.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	switch logType {
	case friendtypes.MarryDevelopLogTypeAll:
		logList := friend.GetFriendService().GetMarryDevelopLogByTime(logTime)
		scMsg := pbutil.BuildSCFriendMarryDevelopLogIncr(logType, logList)
		pl.SendMsg(scMsg)
		break
	case friendtypes.MarryDevelopLogTypeSend:
		logSendList := friendManager.GetMarryDevelopSendLogByTime(logTime)
		scMsg := pbutil.BuildSCFriendMarryDevelopSendLogIncr(logType, logSendList)
		pl.SendMsg(scMsg)
		break
	case friendtypes.MarryDevelopLogTypeRecv:
		logRecvList := friendManager.GetMarryDevelopRecvLogByTime(logTime)
		scMsg := pbutil.BuildSCFriendMarryDevelopRecvLogIncr(logType, logRecvList)
		pl.SendMsg(scMsg)
		break
	}
	return
}
