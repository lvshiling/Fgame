package use

import (
	alliancelogic "fgame/fgame/game/alliance/logic"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	playergoldequip "fgame/fgame/game/goldequip/player"
	playerguaji "fgame/fgame/game/guaji/player"
	guajitypes "fgame/fgame/game/guaji/types"
	inventoryguaji "fgame/fgame/game/inventory/guaji/guaji"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertylogic "fgame/fgame/game/property/logic"
)

func init() {
	inventoryguaji.RegisterGuaJiItemUseHandler(itemtypes.ItemTypeGoldEquip, inventoryguaji.GuaJiItemUseHandlerFunc(handleGoldEquipUse))
}

func handleGoldEquipUse(pl player.Player, index int32, num int32) {
	//参数不对
	if num != 1 {
		return
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	guaJiManager := pl.GetPlayerDataManager(playertypes.PlayerGuaJiManagerType).(*playerguaji.PlayerGuaJiManager)
	itemObj := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, index)
	if itemObj == nil {
		return
	}

	if itemObj.IsEmpty() {
		return
	}

	itemId := itemObj.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return
	}

	if !itemTemplate.IsGoldEquip() {
		return
	}
	equipmentTemplate := itemTemplate.GetGoldEquipTemplate()
	equipmentPower := propertylogic.CulculateForce(equipmentTemplate.GetBattlePropertyMap())
	equipmentSubType := itemTemplate.GetItemSubType().(itemtypes.ItemGoldEquipSubType)
	pos := equipmentSubType.Position()
	goldequipManager := pl.GetPlayerDataManager(playertypes.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	//判断是否已经装备了
	equipmentItem := goldequipManager.GetGoldEquipByPos(pos)

	useFlag := true
	//检查角色
	if itemTemplate.NeedProfession != 0 {
		if itemTemplate.GetRole() != pl.GetRole() {
			goto Recycle
		}
	}
	//检查性别
	if itemTemplate.GetSex() != 0 {
		//性别
		if itemTemplate.GetSex() != pl.GetSex() {
			goto Recycle
		}
	}
	//判断级别
	if itemTemplate.NeedLevel > pl.GetLevel() {
		return
	}

	//判断转数
	if itemTemplate.NeedZhuanShu > pl.GetZhuanSheng() {
		return
	}

	if equipmentItem != nil && !equipmentItem.IsEmpty() {
		currentItemTemplate := item.GetItemService().GetItem(int(equipmentItem.GetItemId()))
		//判断战斗力
		currentEquipmentTemplate := currentItemTemplate.GetGoldEquipTemplate()
		if itemTemplate.NeedZhuanShu < currentItemTemplate.NeedZhuanShu {
			useFlag = false
		} else if itemTemplate.NeedZhuanShu == currentItemTemplate.NeedZhuanShu {
			currentEquipmentPower := propertylogic.CulculateForce(currentEquipmentTemplate.GetBattlePropertyMap())
			if equipmentPower < currentEquipmentPower {
				useFlag = false
			}
		}
	}
	if useFlag {
		//使用装备
		goldequiplogic.HandleUseGoldEquip(pl, index)
		return
	}
Recycle:
	emptySlots := inventoryManager.GetEmptySlots(inventorytypes.BagTypePrim)
	if emptySlots > guaJiManager.GetGlobalValue(guajitypes.GuaJiGlobalTypeBagRemainSlots) {
		return
	}
	//检查仙盟仓库
	flag := alliancelogic.CheckPlayerIfCanSaveAllianceDepot(pl, index, num)
	if flag {
		alliancelogic.HandleSaveAllianceDepot(pl, index, num)
		return
	}
	//检查仓库
	flag = inventorylogic.CheckPlayerIfCanSaveInDepot(pl, index)
	if flag {
		inventorylogic.HandleSaveInDepot(pl, index)
		return
	}
	//分解
	goldequiplogic.HandleGoldEquipEat(pl, 0, []int32{index})
	return
}
