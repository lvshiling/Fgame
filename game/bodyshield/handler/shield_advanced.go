package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	bodyshieldlogic "fgame/fgame/game/bodyshield/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_SHIELD_ADVANCED_TYPE), dispatch.HandlerFunc(handleShieldAdvanced))
}

//处理神盾尖刺进阶信息
func handleShieldAdvanced(s session.Session, msg interface{}) (err error) {
	log.Debug("bodyshield:处理神盾尖刺进阶信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csShieldAdvanced := msg.(*uipb.CSShieldAdvanced)
	autoFlag := csShieldAdvanced.GetAutoFlag()

	err = bodyshieldlogic.HandleShieldAdvanced(tpl, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("bodyshield:处理神盾尖刺进阶信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("bodyshield:处理神盾尖刺进阶完成")
	return nil

}
