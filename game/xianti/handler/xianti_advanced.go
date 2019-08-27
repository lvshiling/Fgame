package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	xiantilogic "fgame/fgame/game/xianti/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_XIANTI_ADVANCED_TYPE), dispatch.HandlerFunc(handleXianTiAdvanced))
}

//处理仙体进阶信息
func handleXianTiAdvanced(s session.Session, msg interface{}) (err error) {
	log.Debug("xianti:处理仙体进阶信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csXianTiAdvanced := msg.(*uipb.CSXiantiAdvanced)
	autoFlag := csXianTiAdvanced.GetAutoFlag()

	err = xiantilogic.HandleXianTiAdvanced(tpl, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("xianti:处理仙体进阶信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("xianti:处理仙体进阶完成")
	return nil
}
