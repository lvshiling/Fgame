package use

import (
	inventoryguaji "fgame/fgame/game/inventory/guaji/guaji"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerskill "fgame/fgame/game/skill/player"
)

func init() {
	inventoryguaji.RegisterGuaJiItemAutoUseHandler(itemtypes.ItemTypeSkill, itemtypes.ItemDefaultSubTypeDefault, inventoryguaji.GuaJiItemAutoUseHandlerFunc(handleGuaJiItemAutoUse))
}

func handleGuaJiItemAutoUse(pl player.Player, index int32, num int32) {
	//参数不对
	if num <= 0 {
		return
	}
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	itemObj := inventoryManager.FindItemByIndex(inventorytypes.BagTypePrim, index)
	if itemObj == nil {
		return
	}
	if itemObj.IsEmpty() {
		return
	}
	itemId := itemObj.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	skillTemplate := itemTemplate.GetSkillTemplate()
	skillManager := pl.GetPlayerDataManager(playertypes.PlayerSkillDataManagerType).(*playerskill.PlayerSkillDataManager)
	if skillTemplate.GetRoleType() != pl.GetRole() {
		return
	}

	skillId := int32(skillTemplate.TemplateId())
	exist := skillManager.IfSkillExist(skillId)
	if exist {
		return
	}

	inventorylogic.UseItemIndex(pl, inventorytypes.BagTypePrim, index, 1, nil, "")

	return
}
