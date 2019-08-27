package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/inventory/logic"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"

	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_INVENTORY_ITEM_USE_TYPE), dispatch.HandlerFunc(handleInventoryItemUse))
}

func handleInventoryItemUse(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理背包使用消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csInventoryItemUse := msg.(*uipb.CSInventoryItemUse)
	index := csInventoryItemUse.GetIndex()
	num := csInventoryItemUse.GetNum()
	args := csInventoryItemUse.GetArgs()

	var chooseIndexList []int32
	chooseInfo := csInventoryItemUse.GetChooseInfo()
	if chooseInfo != nil {
		chooseIndexList = chooseInfo.GetChooseIndexList()
	}
	if num <= 0 || index < 0 {
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"index":    index,
				"num":      num,
			}).Warn("inventory:处理背包使用,参数不对")
		return
	}
	bagType := inventorytypes.BagTypePrim
	bagTypePtr := csInventoryItemUse.BagType
	if bagTypePtr != nil {
		bagType = inventorytypes.BagType(csInventoryItemUse.GetBagType())
		if !bagType.Valid() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"bagType":  bagType,
				}).Warn("inventory:处理背包使用,参数不对")
			playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
			return
		}
	}

	err = useItemIndex(tpl, bagType, index, num, chooseIndexList, args)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": tpl.GetId(),
				"index":    index,
				"num":      num,
			}).Error("inventory:处理背包使用,错误")

		return
	}
	log.WithFields(
		log.Fields{
			"playerId": tpl.GetId(),
			"index":    index,
			"num":      num,
		}).Debug("inventory:处理背包使用消息,成功")

	return
}

func useItemIndex(pl player.Player, bagType inventorytypes.BagType, index int32, num int32, chooseIndexList []int32, args string) (err error) {
	return logic.UseItemIndex(pl, bagType, index, num, chooseIndexList, args)
}
