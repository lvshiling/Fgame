package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/treasurebox/pbutil"
	"fgame/fgame/game/treasurebox/treasurebox"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_TREASUREBOX_LOG_TYPE), dispatch.HandlerFunc(handleTreasureBoxLog))
}

//处理跨服宝箱日志信息
func handleTreasureBoxLog(s session.Session, msg interface{}) (err error) {
	log.Debug("treasurebox:处理获取跨服宝箱日志消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csTreasureBoxLog := msg.(*uipb.CSTreasureBoxLog)
	logTime := csTreasureBoxLog.GetLogTime()

	err = treasureBoxLog(tpl, logTime)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("soul:处理获取跨服宝箱日志消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("soul:处理获取跨服宝箱日志消息完成")
	return nil

}

//获取跨服宝箱日志界面信息的逻辑
func treasureBoxLog(pl player.Player, logTime int64) (err error) {
	if logTime < 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"logTime":  logTime,
		}).Warn("treasurebox:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	boxLogList := treasurebox.GetTreasureBoxService().GetTreasureBoxLogList(logTime)
	scTreasureBoxLog := pbutil.BuildSCTreasureBoxLog(boxLogList)
	pl.SendMsg(scTreasureBoxLog)
	return
}
