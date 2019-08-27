package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	winglogic "fgame/fgame/game/wing/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_WING_ADVANCED_TYPE), dispatch.HandlerFunc(handleWingAdvanced))
}

//处理战翼进阶信息
func handleWingAdvanced(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理战翼进阶信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csWingAdvanced := msg.(*uipb.CSWingAdvanced)
	autoFlag := csWingAdvanced.GetAutoFlag()

	err = wingAdvanced(tpl, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("wing:处理战翼进阶信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("wing:处理战翼进阶完成")
	return nil

}

//战翼进阶的逻辑
func wingAdvanced(pl player.Player, autoFlag bool) (err error) {
	return winglogic.HandleWingAdvanced(pl, autoFlag)
}
