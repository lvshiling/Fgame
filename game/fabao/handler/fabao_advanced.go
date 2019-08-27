package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	fabaologic "fgame/fgame/game/fabao/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_FABAO_ADVANCED_TYPE), dispatch.HandlerFunc(handleFaBaoAdvanced))
}

//处理法宝进阶信息
func handleFaBaoAdvanced(s session.Session, msg interface{}) (err error) {
	log.Debug("fabao:处理法宝进阶信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csFaBaoAdvanced := msg.(*uipb.CSFaBaoAdvanced)
	autoFlag := csFaBaoAdvanced.GetAutoFlag()

	err = fabaologic.HandleFaBaoAdvanced(tpl, autoFlag)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("fabao:处理法宝进阶信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("fabao:处理法宝进阶完成")
	return nil

}
