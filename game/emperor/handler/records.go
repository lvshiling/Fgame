package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/emperor/emperor"
	"fgame/fgame/game/emperor/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_EMPEROR_RECORD_TYPE), dispatch.HandlerFunc(handleEmperorRecord))
}

//处理抢夺记录信息
func handleEmperorRecord(s session.Session, msg interface{}) (err error) {
	log.Debug("emperor:处理抢夺记录信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csEmperorRecords := msg.(*uipb.CSEmperorRecords)
	logTime := csEmperorRecords.GetLogTime()

	err = emperorRecord(tpl, logTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("emperor:处理抢夺记录信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("emperor:处理抢夺记录信息完成")
	return nil
}

//处理抢夺记录界面信息逻辑
func emperorRecord(pl player.Player, logTime int64) (err error) {
	if logTime < 0 {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
			"logTime":  logTime,
		}).Warn("emperor:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	robList := emperor.GetEmperorService().GetEmperorRobListByLogTime(logTime)
	scEmperorRecordsGet := pbuitl.BuildSCEmperorRecordsGet(robList)
	pl.SendMsg(scEmperorRecordsGet)
	return
}
