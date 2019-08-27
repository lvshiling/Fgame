package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/welfare/pbutil"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_DREW_LOG_INCR_TYPE), dispatch.HandlerFunc(handleDrewLogIncr))
}

//处理抽奖日志请求
func handleDrewLogIncr(s session.Session, msg interface{}) (err error) {
	log.Debug("drew:处理获取抽奖日志请求")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSOpenActivityDrewLogIncr)
	groupId := csMsg.GetGroupId()
	logTime := csMsg.GetLogTime()

	err = drewLogIncr(tpl, groupId, logTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"logTime":  logTime,
				"error":    err,
			}).Error("drew:处理获取抽奖日志请求,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"logTime":  logTime,
		}).Debug("drew:处理获取抽奖日志请求完成")
	return nil

}

//获取抽奖界面信息逻辑
func drewLogIncr(pl player.Player, groupId int32, logTime int64) (err error) {
	logList := welfare.GetWelfareService().GetDrewLogByTime(groupId, logTime)
	if len(logList) < 1 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"logTime":  logTime,
			}).Info("drew:处理获取抽奖日志请求,日志增量列表为空")
	}

	scMsg := pbutil.BuildSCOpenActivityDrewLogIncr(logList, groupId)
	pl.SendMsg(scMsg)
	return
}
