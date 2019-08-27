package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	lingyulogic "fgame/fgame/game/lingyu/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGYU_ADVANCED_TYPE), dispatch.HandlerFunc(handleLingyuAdvanced))
}

//处理领域进阶信息
func handleLingyuAdvanced(s session.Session, msg interface{}) (err error) {
	log.Debug("lingyu:处理领域进阶信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingyuAdvanced := msg.(*uipb.CSLingyuAdvanced)
	autoFlag := csLingyuAdvanced.GetAutoFlag()

	err = lingyulogic.HandleLingyuAdvanced(tpl, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("lingyu:处理领域进阶信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("lingyu:处理领域进阶完成")
	return nil

}
