package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/inventory/pbutil"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_INVENTORY_MERGE_TYPE), dispatch.HandlerFunc(handleInventoryMerge))
}

//处理合并
func handleInventoryMerge(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理合并背包")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csInventoryMerge := msg.(*uipb.CSInventoryMerge)
	bagType := inventorytypes.BagTypePrim
	bagTypePtr := csInventoryMerge.BagType
	if bagTypePtr != nil {
		bagType = inventorytypes.BagType(csInventoryMerge.GetBagType())
		if !bagType.Valid() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"bagType":  bagType,
				}).Warn("inventory:处理合并背包,参数不对")
			playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
			return
		}
	}
	err = merge(tpl, bagType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("inventory:处理合并背包,错误")

		return
	}
	log.Debug("inventory:处理合并背包,完成")
	return
}

//合并
func merge(pl player.Player, bagType inventorytypes.BagType) (err error) {
	//TODO 加cd
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	manager.Merge(bagType)
	scInventoryMerge := pbutil.BuildSCInventoryMerge(bagType, manager.GetBagAll(bagType))
	pl.SendMsg(scInventoryMerge)
	return
}
