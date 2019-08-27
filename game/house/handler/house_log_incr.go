package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/house/house"
	"fgame/fgame/game/house/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_HOUSE_LOG_INCR_TYPE), dispatch.HandlerFunc(handleHouseLogIncr))
}

//处理房子日志请求
func handleHouseLogIncr(s session.Session, msg interface{}) (err error) {
	log.Debug("house:处理获取房子日志请求")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSHouseLogIncr)
	logTime := csMsg.GetLogTime()

	err = houseLogIncr(tpl, logTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"logTime":  logTime,
				"error":    err,
			}).Error("house:处理获取房子日志请求,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"logTime":  logTime,
		}).Debug("house:处理获取房子日志请求完成")
	return nil

}

//获取房子日志列表逻辑
// 客户端定时请求，更新日志列表
func houseLogIncr(pl player.Player, logTime int64) (err error) {
	logList := house.GetHouseService().GetLogByTime(logTime)
	if len(logList) < 1 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"logTime":  logTime,
			}).Warn("house:处理获取房子日志请求,日志增量列表为空")
		return
	}

	scMsg := pbutil.BuildSCHouseLogIncr(logList)
	pl.SendMsg(scMsg)
	return
}
