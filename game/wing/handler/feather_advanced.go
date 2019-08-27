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
	processor.Register(codec.MessageType(uipb.MessageType_CS_FEATHER_ADVANCED_TYPE), dispatch.HandlerFunc(handleFeatherAdvanced))
}

//处理护体仙羽进阶信息
func handleFeatherAdvanced(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理护体仙羽进阶信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csFeatherAdvanced := msg.(*uipb.CSFeatherAdvanced)
	autoFlag := csFeatherAdvanced.GetAutoFlag()

	err = winglogic.HandleFeatherAdvanced(tpl, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("wing:处理护体仙羽进阶信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("wing:处理护体仙羽进阶完成")
	return nil

}
