package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SAVE_IN_DEPOT_TYPE), dispatch.HandlerFunc(handleSaveInDepot))
}

//处理存入仓库
func handleSaveInDepot(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理存入仓库")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csSaveInDepot := msg.(*uipb.CSSaveInDepot)
	itemIndex := csSaveInDepot.GetIndex()

	err = saveInDepot(tpl, itemIndex)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"itemIndex": itemIndex,
				"error":     err,
			}).Error("inventory:处理存入仓库,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"itemIndex": itemIndex,
		}).Debug("inventory:处理存入仓库,完成")
	return nil
}

//存入仓库
func saveInDepot(pl player.Player, itemIndex int32) (err error) {
	return inventorylogic.HandleSaveInDepot(pl, itemIndex)
}
