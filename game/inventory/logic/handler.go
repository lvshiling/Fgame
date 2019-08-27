package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/inventory/pbutil"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
)

//存入仓库
func CheckPlayerIfCanSaveInDepot(pl player.Player, itemIndex int32) (flag bool) {
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	itemObj := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, itemIndex)
	//物品不存在
	if itemObj == nil || itemObj.IsEmpty() {
		return
	}

	itemId := itemObj.ItemId
	itemNum := itemObj.Num
	level := itemObj.Level
	bind := itemObj.BindType
	propertyData := itemObj.PropertyData

	itemTemplate := item.GetItemService().GetItem(int(itemId))
	//物品是否能够存入仓库
	if itemTemplate.Storage == 0 {
		return
	}

	//仓库空间是否足够
	return inventoryManager.HasEnoughDepotSlotWithProperty(itemId, itemNum, level, bind, propertyData.GetExpireType(), propertyData.GetItemGetTime(), propertyData.GetExpireTime())

}

//存入仓库
func HandleSaveInDepot(pl player.Player, itemIndex int32) (err error) {
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	itemObj := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, itemIndex)
	//物品不存在
	if itemObj == nil || itemObj.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"itemIndex": itemIndex,
			}).Warn("inventory:存放仓库失败，该物品不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	itemId := itemObj.ItemId
	itemNum := itemObj.Num
	level := itemObj.Level
	bind := itemObj.BindType
	propertyData := itemObj.PropertyData

	itemTemplate := item.GetItemService().GetItem(int(itemId))
	//物品是否能够存入仓库
	if itemTemplate.Storage == 0 {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"itemIndex": itemIndex,
				"itemId":    itemId,
				"itemNum":   itemNum,
			}).Warn("inventory:存放仓库失败，该物品不允许放入仓库")
		playerlogic.SendSystemMessage(pl, lang.InventoryDepotNotAllowStore)
		return
	}

	//仓库空间是否足够
	flag := inventoryManager.HasEnoughDepotSlotWithProperty(itemId, itemNum, level, bind, propertyData.GetExpireType(), propertyData.GetItemGetTime(), propertyData.GetExpireTime())
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"itemIndex": itemIndex,
				"itemId":    itemId,
				"itemNum":   itemNum,
			}).Warn("inventory:存放仓库失败，仓库空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryDepotSlotNoEnough)
		return
	}

	itemUseReason := commonlog.InventoryLogReasonSaveInDepot
	flag, _ = inventoryManager.RemoveIndex(inventorytypes.BagTypePrim, itemIndex, itemNum, itemUseReason, itemUseReason.String())
	if !flag {
		panic("inventory: 背包移除物品应该成功")
	}

	flag = inventoryManager.AddItemInDepot(itemId, itemNum, level, bind, propertyData)
	if !flag {
		panic("inventory:存入仓库应该成功")
	}

	SnapInventoryChanged(pl)
	SnapDepotChanged(pl)

	itemChangedList := inventoryManager.GetDepotChangedSlotAndReset()
	scSaveInDepot := pbutil.BuildSCSaveInDepot(itemChangedList)
	pl.SendMsg(scSaveInDepot)
	return
}

//获取可以购买的槽位数量
func GetNumOfRemainBuySlots(pl player.Player) (num int32) {
	singleSlotGold := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeGoldForSingleSlot)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	canBuyNum := propertyManager.GetGold() / int64(singleSlotGold)
	if canBuyNum <= 0 {
		return 0
	}
	//超过32位
	if canBuyNum >= math.MaxInt32 {
		return math.MaxInt32
	}
	num = int32(canBuyNum)
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	remainBuySlots := manager.NumOfRemainBuySlots()
	if num >= remainBuySlots {
		return remainBuySlots
	} else {
		return num
	}

}

//购买槽位
func HandleBuySlots(pl player.Player, buyNum int32) (err error) {
	singleSlotGold := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeGoldForSingleSlot)
	openGold := int64(singleSlotGold * buyNum)
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	flag := propertyManager.HasEnoughGold(int64(openGold), true)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":      pl.GetId(),
				"buyNum":        buyNum,
				"totalNeedGold": openGold,
			}).Warn("inventory:处理获取背包购买槽位,元宝不足")

		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	reasonText := commonlog.GoldLogReasonBuySlots.String()
	flag = propertyManager.CostGold(openGold, true, commonlog.GoldLogReasonBuySlots, reasonText)
	if !flag {
		panic(fmt.Errorf("inventory:花费元宝应该成功"))
	}
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	flag = manager.IfCanAddSlots(buyNum)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("inventory:处理获取背包购买槽位,不能购买槽位了")

		playerlogic.SendSystemMessage(pl, lang.InventoryCanNotAddSlot)
		return
	}
	flag = manager.AddSlots(buyNum)
	if !flag {
		panic(fmt.Errorf("inventory:add slots should be ok"))
	}
	propertylogic.SnapChangedProperty(pl)

	slotsNum := manager.GetSlots()
	inventoryBuySlots := pbutil.BuildSCInventoryBuySlots(slotsNum)
	pl.SendMsg(inventoryBuySlots)

	return nil
}
