package logic

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"fgame/fgame/game/tulongequip/pbutil"
	playertulongequip "fgame/fgame/game/tulongequip/player"
	tulongequiptypes "fgame/fgame/game/tulongequip/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

//推送屠龙装备改变
func SnapInventoryTuLongEquipChanged(pl player.Player) (err error) {
	tulongequipManager := pl.GetPlayerDataManager(playertypes.PlayerTuLongEquipDataManagerType).(*playertulongequip.PlayerTuLongEquipDataManager)
	slotChangedMap := tulongequipManager.GetChangedEquipmentSlotAndResetMap()
	if len(slotChangedMap) <= 0 {
		return
	}
	scMsg := pbutil.BuildSCTuLongEquipSlotChanged(slotChangedMap)
	pl.SendMsg(scMsg)
	return
}

//屠龙装备属性
func TuLongEquipPropertyChanged(pl player.Player) (err error) {
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeTuLongEquip.Mask())
	return
}

//使用屠龙装备
func HandleUseTuLongEquip(pl player.Player, index int32, suitType tulongequiptypes.TuLongSuitType) (err error) {
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	tulongequipManager := pl.GetPlayerDataManager(playertypes.PlayerTuLongEquipDataManagerType).(*playertulongequip.PlayerTuLongEquipDataManager)
	it := inventoryManager.FindItemByIndex(inventorytypes.BagTypeTuLongEquip, index)
	//物品不存在
	if it == nil || it.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("inventory:使用屠龙装备,物品不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}
	itemId := it.ItemId
	propertyData := it.PropertyData
	bind := it.BindType
	//判断物品是否可以装备
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if !itemTemplate.IsTuLongEquip() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"itemType": itemTemplate.GetItemType(),
			}).Warn("inventory:使用屠龙装备,此物品不是屠龙装备")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemCanNotEquip)
		return
	}

	if itemTemplate.GetSex() != 0 {
		//性别
		if itemTemplate.GetSex() != pl.GetSex() {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"index":    index,
				}).Warn("inventory:使用屠龙装备,性别不符")
			playerlogic.SendSystemMessage(pl, lang.PlayerSexWrong)
			return
		}
	}

	//判断级别
	if itemTemplate.NeedLevel > pl.GetLevel() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
			}).Warn("inventory:使用屠龙装备,等级不够")
		playerlogic.SendSystemMessage(pl, lang.PlayerLevelTooLow)
		return
	}

	//判断是否已经装备

	pos := itemTemplate.GetTuLongEquipTemplate().GetPosType()
	if !pos.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"index":    index,
				"pos":      pos,
			}).Warn("inventory:使用屠龙装备,没有可装备的位置")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	equipmentItem := tulongequipManager.GetTuLongEquipByPos(suitType, pos)
	if equipmentItem != nil && !equipmentItem.IsEmpty() {
		flag := takeOffInternal(pl, suitType, pos)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"pos":      pos,
				}).Warn("inventory:使用屠龙装备,穿戴前卸下装备错误")
			return
		}
	}

	//移除物品
	reasonText := commonlog.InventoryLogReasonPutOn.String()
	flag, _ := inventoryManager.RemoveIndex(inventorytypes.BagTypeTuLongEquip, index, 1, commonlog.InventoryLogReasonPutOn, reasonText)
	if !flag {
		panic("inventory:移除物品应该是可以的")
	}

	flag = tulongequipManager.PutOn(suitType, pos, itemId, bind, propertyData)
	if !flag {
		panic(fmt.Errorf("inventory:穿上位置 [%s]应该是可以的", pos.String()))
	}

	//同步改变
	TuLongEquipPropertyChanged(pl)
	SnapInventoryTuLongEquipChanged(pl)
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	scMsg := pbutil.BuildSCTuLongTakeOffEquip(index, int32(suitType))
	pl.SendMsg(scMsg)
	return
}

//脱下
func HandleTakeOff(pl player.Player, suitType tulongequiptypes.TuLongSuitType, pos inventorytypes.BodyPositionType) (err error) {
	flag := takeOffInternal(pl, suitType, pos)
	if !flag {
		return
	}
	//同步改变
	TuLongEquipPropertyChanged(pl)
	SnapInventoryTuLongEquipChanged(pl)
	inventorylogic.SnapInventoryChanged(pl)
	propertylogic.SnapChangedProperty(pl)

	//脱下成功
	scMsg := pbutil.BuildSCTuLongTakeOffEquip(int32(suitType), int32(pos))
	pl.SendMsg(scMsg)

	return
}

func takeOffInternal(pl player.Player, suitType tulongequiptypes.TuLongSuitType, pos inventorytypes.BodyPositionType) (flag bool) {
	tulongequipManager := pl.GetPlayerDataManager(playertypes.PlayerTuLongEquipDataManagerType).(*playertulongequip.PlayerTuLongEquipDataManager)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	//没有东西
	slot := tulongequipManager.GetTuLongEquipByPos(suitType, pos)
	if slot == nil || slot.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pos":      pos,
			}).Warn("tulongequip:脱下屠龙装备,屠龙装备不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryEquipCanNotTakeOff)
		return
	}
	level := int32(0)
	num := int32(1)
	bind := slot.GetBindType()
	propertyData := slot.GetPropertyData()

	//背包空间
	if !inventoryManager.HasEnoughSlotItemLevel(slot.GetItemId(), num, level, bind) {
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	itemId := tulongequipManager.TakeOff(suitType, pos)
	if itemId == 0 {
		panic(fmt.Errorf("tulongequip:take off should more than 0"))
	}

	//添加物品
	reasonText := commonlog.InventoryLogReasonTakeOff.String()
	flag = inventoryManager.AddItemLevelWithPropertyData(itemId, num, level, bind, propertyData, commonlog.InventoryLogReasonTakeOff, reasonText)
	if !flag {
		panic(fmt.Errorf("tulongequip:add item should be success"))
	}
	return
}
