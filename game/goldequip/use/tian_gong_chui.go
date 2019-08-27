package use

import (
	"fgame/fgame/common/lang"
	"fgame/fgame/core/utils"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	playergoldequip "fgame/fgame/game/goldequip/player"
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
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeGoldEquipStrengthen, itemtypes.ItemGoldEquipStrengthenSubTypeChuiZi, playerinventory.ItemUseHandleFunc(handlerGoldEquipStrengthen))
}

//元神金装强化道具
// args:param1:1背包2身体；param2:强化位置
func handlerGoldEquipStrengthen(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
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
			}).Warn("goldequip:强化升级失败,参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	positionType := paramsIntArr[0]
	targetIndex := paramsIntArr[1]
	if positionType == 1 {
		flag, err = upLevelBagEquip(pl, itemId, int32(targetIndex))
	} else {
		flag, err = upLevelBodyEquip(pl, itemId, int32(targetIndex))
	}

	return
}

func upLevelBagEquip(pl player.Player, itemId int32, targetIndex int32) (flag bool, err error) {
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	targetIt := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, targetIndex)
	//物品不存在
	if targetIt == nil || targetIt.IsEmpty() {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetIndex,
			}).Warn("goldequip:强化升级失败,强化目标不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	//能否被强化
	targetItemTemplate := item.GetItemService().GetItem(int(targetIt.ItemId))
	goldEquipTemplate := targetItemTemplate.GetGoldEquipTemplate()
	if goldEquipTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"targetItemId": targetIt.ItemId,
				"targetIndex":  targetIndex,
			}).Warn("goldequip:强化升级失败,强化模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if goldEquipTemplate.GoldequipStrenId == 0 {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetIndex,
			}).Warn("goldequip:强化升级失败,该金装无法被强化")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipStrengthenNotAllow)
		return
	}

	//元神强化等级是否超过
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	toUplevel := itemTemplate.TypeFlag1
	curLevel := targetIt.Level
	if curLevel >= toUplevel {
		log.WithFields(
			log.Fields{
				"playerId":    pl.GetId(),
				"targetIndex": targetIndex,
			}).Warn("goldequip:强化升级失败,当前元神装备强化等级已超过")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipLevelToHigh)
		return
	}

	//成功
	flag = inventoryManager.UpdateGoldEquipLevelUseItem(targetIndex, itemId)
	if !flag {
		panic(fmt.Errorf("goldequip: 强化升级应该成功"))
	}

	//同步改变
	inventorylogic.SnapInventoryChanged(pl)

	flag = true
	return
}

func upLevelBodyEquip(pl player.Player, itemId int32, targetSlotId int32) (flag bool, err error) {
	posType := inventorytypes.BodyPositionType(targetSlotId)
	if !posType.Valid() {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"targetSlotId": targetSlotId,
				"error":        err,
			}).Warn("goldequip:处理装备槽金装强化,位置错误")
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
			}).Warn("goldequip:强化升级失败,强化目标不存在")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoExist)
		return
	}

	//能否被强化
	targetItemTemp := item.GetItemService().GetItem(int(targetIt.GetItemId()))
	goldEquipTemplate := targetItemTemp.GetGoldEquipTemplate()
	if goldEquipTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"targetItemId": targetIt.GetItemId(),
				"targetIndex":  targetSlotId,
			}).Warn("goldequip:强化升级失败,强化模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	if goldEquipTemplate.GoldequipStrenId == 0 {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"targetSlotId": targetSlotId,
			}).Warn("goldequip:强化升级失败,该金装无法被强化")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipStrengthenNotAllow)
		return
	}

	//元神强化等级是否超过
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	toUplevel := itemTemplate.TypeFlag1
	curLevel := targetIt.GetLevel()
	if curLevel >= toUplevel {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"targetSlotId": targetSlotId,
			}).Warn("goldequip:强化升级失败,当前元神装备强化等级已超过")
		playerlogic.SendSystemMessage(pl, lang.GoldEquipLevelToHigh)
		return
	}

	equipBag.UpdateGoldEquipLevelUseItem(posType, itemId)

	// 同步
	goldequiplogic.GoldEquipPropertyChanged(pl)
	goldequiplogic.SnapInventoryGoldEquipChanged(pl)

	flag = true
	return
}
