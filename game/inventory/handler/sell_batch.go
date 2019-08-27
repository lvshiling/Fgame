package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/inventory/pbutil"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_INVENTORY_ITEM_SELL_BATCH_TYPE), dispatch.HandlerFunc(handleInventorySellBatch))
}

//处理批量出售物品
func handleInventorySellBatch(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理出售物品")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csInventoryItemSellBatch := msg.(*uipb.CSInventoryItemSellBatch)
	indexList := csInventoryItemSellBatch.GetIndexList()

	bagType := inventorytypes.BagTypePrim
	bagTypePtr := csInventoryItemSellBatch.BagType
	if bagTypePtr != nil {
		bagType = inventorytypes.BagType(csInventoryItemSellBatch.GetBagType())
		if !bagType.Valid() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"bagType":  bagType,
				}).Warn("inventory:处理出售物品,参数不对")
			playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
			return
		}
	}

	err = sellBatch(tpl, bagType, indexList)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"indexList": indexList,
				"error":     err,
			}).Error("inventory:处理出售物品,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"indexList": indexList,
		}).Debug("inventory:处理出售物品,完成")
	return nil
}

//批量出售
func sellBatch(pl player.Player, bagType inventorytypes.BagType, indexList []int32) (err error) {
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	if coreutils.IfRepeatElementInt32(indexList) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"indexList": indexList,
			}).Warn("inventory:处理出售物品,索引重复")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	totalSilverNum := int64(0)
	for _, index := range indexList {
		it := inventoryManager.FindItemByIndex(bagType, index)
		if it == nil {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"bagType":   bagType,
					"indexList": indexList,
					"sellIndex": index,
				}).Warn("inventory:处理出售物品,索引不存在")
			playerlogic.SendSystemMessage(pl, lang.InventoryCanNotSell)
			return
		}

		itemTemplate := item.GetItemService().GetItem(int(it.ItemId))
		if !inventoryManager.IfCanSell(bagType, index, it.Num) {
			log.WithFields(
				log.Fields{
					"playerId":  pl.GetId(),
					"bagType":   bagType,
					"indexList": indexList,
					"sellIndex": index,
				}).Warn("inventory:处理出售物品,不可以出售")
			playerlogic.SendSystemMessage(pl, lang.InventoryCanNotSell)
			return
		}

		num := int64(it.Num)
		silverNum := int64(math.Ceil(float64(itemTemplate.SaleRate) / common.MAX_RATE * float64(itemTemplate.BuySilver)))
		totalSilverNum += num * silverNum

	}

	//移除物品
	itemUseReason := commonlog.InventoryLogReasonSell
	inventoryManager.BatchRemoveIndex(bagType, indexList, itemUseReason, itemUseReason.String())

	//加钱
	silverReasonText := fmt.Sprintf(commonlog.SilverLogReasonItemSellBatch.String(), indexList)
	propertyManager.AddSilver(totalSilverNum, commonlog.SilverLogReasonItemSellBatch, silverReasonText)

	//物品改变
	inventorylogic.SnapInventoryChanged(pl)
	//属性变化
	propertylogic.SnapChangedProperty(pl)

	inventoryBuySlots := pbutil.SCInventoryItemSellBatch(totalSilverNum)
	pl.SendMsg(inventoryBuySlots)
	return nil
}
