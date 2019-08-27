package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/outlandboss/outlandboss"
	"fgame/fgame/game/outlandboss/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OUTLAND_BOSS_DROP_RECORDS_INCR_TYPE), dispatch.HandlerFunc(handleDropRecordsIncr))
}

//处理掉落记录增量请求
func handleDropRecordsIncr(s session.Session, msg interface{}) (err error) {
	log.Debug("outlandboss:处理获取掉落记录请求")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csDropRecordsIncr := msg.(*uipb.CSOutlandBossDropRecordsIncr)
	logTime := csDropRecordsIncr.GetRecordsTime()

	err = dropRecordsIncr(tpl, logTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"logTime":  logTime,
				"error":    err,
			}).Error("outlandboss:处理获取掉落记录请求,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"logTime":  logTime,
		}).Debug("outlandboss:处理获取掉落记录请求完成")
	return nil

}

//获取掉落记录增量逻辑
func dropRecordsIncr(pl player.Player, logTime int64) (err error) {
	logList := outlandboss.GetOutlandBossService().GetDropRecordsList()
	index, _ := outlandboss.GetOutlandBossService().GetDropRecordsByTime(logTime)
	if index == -1 && logTime != 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"logTime":  logTime,
			}).Warn("outlandboss:处理获取掉落记录请求,时间错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	var newLogList []*outlandboss.OutlandBossDropRecordsObject
	if len(logList) != 0 {
		newLogList = logList[index+1:]
	}

	scDropRecordsIncr := pbutil.BuildSCOutlandBossDropRecordsIncr(newLogList)
	pl.SendMsg(scDropRecordsIncr)
	return
}
