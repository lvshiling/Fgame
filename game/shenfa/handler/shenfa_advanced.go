package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	shenfalogic "fgame/fgame/game/shenfa/logic"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHENFA_ADVANCED_TYPE), dispatch.HandlerFunc(handleShenfaAdvanced))
}

//处理身法进阶信息
func handleShenfaAdvanced(s session.Session, msg interface{}) (err error) {
	log.Debug("shenfa:处理身法进阶信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csShenfaAdvanced := msg.(*uipb.CSShenfaAdvanced)
	autoFlag := csShenfaAdvanced.GetAutoFlag()

	err = shenfalogic.HandleShenfaAdvanced(tpl, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("shenfa:处理身法进阶信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("shenfa:处理身法进阶完成")
	return nil

}
