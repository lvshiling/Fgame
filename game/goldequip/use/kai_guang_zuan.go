package use

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/utils"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	playergoldequip "fgame/fgame/game/goldequip/player"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"

	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeGoldEquipStrengthen, itemtypes.ItemGoldEquipStrengthenSubTypeKaiGuangZuan, playerinventory.ItemUseHandleFunc(handlerGoldEquipKaiGuang))
}

//元神金装开光道具
// args:param1:1背包2身体；param2:开光位置
func handlerGoldEquipKaiGuang(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	paramsIntArr, err := utils.SplitAsIntArray(args)
	if err != nil {
		return
	}

	if len(paramsIntArr) != 2 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"args":     args,
			}).Warn("goldequip:开光升级失败,参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	positionType := paramsIntArr[0]
	targetIndex := paramsIntArr[1]
	if positionType == 1 {
		flag, err = kaiGuangBagEquip(pl, itemId, int32(targetIndex))
	} else {
		flag, err = kaiGuangBodyEquip(pl, itemId, int32(targetIndex))
	}

	return
}

func kaiGuangBagEquip(pl player.Player, itemId int32, targetIndex int32) (flag bool, err error) {
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	targetIt := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, targetIndex)
	//物品不存在
	if targetIt == nil || targetIt.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetIndex,
			}).Warn("goldequip:开光升级失败,开光目标不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	//能否被开光
	targetItemTemplate := item.GetItemService().GetItem(int(targetIt.ItemId))
	goldEquipTemplate := targetItemTemplate.GetGoldEquipTemplate()
	if goldEquipTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"targetItemId": targetIt.ItemId,
				"targetIndex":  targetIndex,
			}).Warn("goldequip:开光升级失败,开光模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if goldEquipTemplate.GoldeuipOpenlightId == 0 {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetIndex,
			}).Warn("goldequip:开光升级失败,该金装无法被开光")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipStrengthenNotAllow)
		return
	}

	//元神开光等级是否超过
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	toUplevel := itemTemplate.TypeFlag1
	equipProperty, ok := targetIt.PropertyData.(*goldequiptypes.GoldEquipPropertyData)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetIndex,
			}).Warn("goldequip:开光升级失败,物品属性转换错误，不是元神金装")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipEquipmentNotGoldEquip)
		return
	}
	if equipProperty.OpenLightLevel >= toUplevel {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetIndex,
				"curLevel":    equipProperty.OpenLightLevel,
				"toUplevel":   toUplevel,
			}).Warn("goldequip:开光升级失败,当前元神装备开光等级已超过")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipLevelToHigh)
		return
	}

	//成功
	flag = inventoryManager.UpdateGoldEquipOpenLightUseItem(targetIndex, itemId)
	if !flag {
		panic(fmt.Errorf("goldequip: 开光升级应该成功"))
	}

	//同步改变
	inventorylogic.SnapInventoryChanged(pl)

	flag = true
	return
}

func kaiGuangBodyEquip(pl player.Player, itemId int32, targetSlotId int32) (flag bool, err error) {
	posType := inventorytypes.BodyPositionType(targetSlotId)
	if !posType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"targetSlotId": targetSlotId,
				"error":        err,
			}).Warn("goldequip:处理装备槽金装开光,位置错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	goldequipManager := pl.GetPlayerDataManager(playertypes.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	equipBag := goldequipManager.GetGoldEquipBag()
	targetIt := equipBag.GetByPosition(posType)

	//物品不存在
	if targetIt == nil || targetIt.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"targetSlotId": targetSlotId,
			}).Warn("goldequip:开光升级失败,开光目标不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	//能否被开光
	targetItemTemp := item.GetItemService().GetItem(int(targetIt.GetItemId()))
	goldEquipTemplate := targetItemTemp.GetGoldEquipTemplate()
	if goldEquipTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"targetItemId": targetIt.GetItemId(),
				"targetIndex":  targetSlotId,
			}).Warn("goldequip:开光升级失败,开光模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if goldEquipTemplate.GoldeuipOpenlightId == 0 {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"targetSlotId": targetSlotId,
			}).Warn("goldequip:开光升级失败,该金装无法被开光")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipStrengthenNotAllow)
		return
	}

	//元神开光等级是否超过
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	toUplevel := itemTemplate.TypeFlag1
	equipProperty, ok := targetIt.GetPropertyData().(*goldequiptypes.GoldEquipPropertyData)
	if !ok {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"targetSlotId": targetSlotId,
			}).Warn("goldequip:开光升级失败,物品属性转换错误，不是元神金装")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipEquipmentNotGoldEquip)
		return
	}
	if equipProperty.OpenLightLevel >= toUplevel {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"targetSlotId": targetSlotId,
				"curLevel":     equipProperty.OpenLightLevel,
				"toUplevel":    toUplevel,
			}).Warn("goldequip:开光升级失败,当前元神装备开光等级已超过")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipLevelToHigh)
		return
	}

	flag = equipBag.UpdateGoldEquipOpenLightUseItem(posType, itemId)
	if !flag {
		panic(fmt.Errorf("goldequip: 开光升级应该成功"))
	}

	// 同步
	goldequiplogic.GoldEquipPropertyChanged(pl)
	goldequiplogic.SnapInventoryGoldEquipChanged(pl)

	flag = true
	return
}
