package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/charge/charge"
	"fgame/fgame/game/charge/pbutil"
	playercharge "fgame/fgame/game/charge/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_NEW_FIRST_CHARGE_RECORD), dispatch.HandlerFunc(handleNewFirstChargeInfo))
}

func handleNewFirstChargeInfo(s session.Session, msg interface{}) (err error) {
	log.Debug("charge:处理获取新首充记录信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = newFirstChargeInfo(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("charge:获取新首充记录信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("charge:处理获取新首充记录信息,完成")

	return
}

func newFirstChargeInfo(pl player.Player) (err error) {
	chargeManager := pl.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
	startTime, duration := charge.GetChargeService().GetNewFirstChargeTime()
	info := chargeManager.GetNewFirstChargeRecordInfo(startTime)
	reocrd := info.GetRecord()
	scNewFirstChargeRecord := pbutil.BuildSCNewFirstChargeRecord(startTime, duration, reocrd)
	pl.SendMsg(scNewFirstChargeRecord)
	return
}
