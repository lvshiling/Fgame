package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/inventory/logic"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_INVENTORY_ITEM_SELL_TYPE), dispatch.HandlerFunc(handleInventorySell))
}

//处理出售物品
func handleInventorySell(s session.Session, msg interface{}) (err error) {
	log.Debug("inventory:处理出售物品")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csInventoryItemSell := msg.(*uipb.CSInventoryItemSell)
	index := csInventoryItemSell.GetIndex()
	num := csInventoryItemSell.GetNum()

	bagType := inventorytypes.BagTypePrim
	bagTypePtr := csInventoryItemSell.BagType
	if bagTypePtr != nil {
		bagType = inventorytypes.BagType(csInventoryItemSell.GetBagType())
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
	if index < 0 || num <= 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"num":      num,
			}).Warn("inventory:处理出售物品,参数不对")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}

	err = sell(tpl, bagType, index, num)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"num":      num,
				"error":    err,
			}).Error("inventory:处理出售物品,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"index":    index,
			"num":      num,
		}).Debug("inventory:处理出售物品,完成")
	return nil
}

//出售
func sell(pl player.Player, bagType inventorytypes.BagType, index int32, num int32) (err error) {
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	if !inventoryManager.IfCanSell(bagType, index, num) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"bagType":  bagType,
				"index":    index,
				"num":      num,
			}).Warn("inventory:处理出售物品,不可以出售")
		playerlogic.SendSystemMessage(pl, lang.InventoryCanNotSell)
		return
	}
	it := inventoryManager.FindItemByIndex(bagType, index)
	itemTemplate := item.GetItemService().GetItem(int(it.ItemId))

	reasonText := commonlog.InventoryLogReasonSell.String()
	inventoryManager.RemoveIndex(bagType, index, num, commonlog.InventoryLogReasonSell, reasonText)

	silverNum := int64(math.Ceil(float64(itemTemplate.SaleRate) / common.MAX_RATE * float64(itemTemplate.BuySilver)))
	totalSilverNum := int64(num) * silverNum
	//加钱
	silverReasonText := fmt.Sprintf(commonlog.SilverLogReasonItemSell.String(), itemTemplate.TemplateId(), num)
	propertyManager.AddSilver(totalSilverNum, commonlog.SilverLogReasonItemSell, silverReasonText)
	//物品改变
	logic.SnapInventoryChanged(pl)

	//属性变化
	propertylogic.SnapChangedProperty(pl)

	inventoryBuySlots := pbutil.BuildSCInventoryItemSell(bagType, index, num)
	pl.SendMsg(inventoryBuySlots)

	return nil
}
