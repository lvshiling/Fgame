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
	processor.Register(codec.MessageType(uipb.MessageType_CS_BODYSHIELD_ADVANCED_TYPE), dispatch.HandlerFunc(handleBodyShieldAdvanced))
}

//处理神盾尖刺进阶信息
func handleBodyShieldAdvanced(s session.Session, msg interface{}) (err error) {
	log.Debug("bodyshield:处理获取神盾尖刺进阶消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csBodyShieldAdvanced := msg.(*uipb.CSBodyShieldAdvanced)
	autoFlag := csBodyShieldAdvanced.GetAutoFlag()

	err = bodyShieldAdvanced(tpl, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"autoFlag": autoFlag,
				"error":    err,
			}).Error("bodyshield:处理获取神盾尖刺进阶消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"autoFlag": autoFlag,
		}).Debug("bodyshield:处理获取神盾尖刺进阶消息完成")
	return nil

}

//神盾尖刺进阶
func bodyShieldAdvanced(pl player.Player, autoFlag bool) (err error) {
	return bodyshieldlogic.HandleBodyShieldAdvanced(pl, autoFlag)
}
