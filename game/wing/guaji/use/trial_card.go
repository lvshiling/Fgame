package use

import (
	inventoryguaji "fgame/fgame/game/inventory/guaji/guaji"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerwing "fgame/fgame/game/wing/player"
)

func init() {
	inventoryguaji.RegisterGuaJiItemAutoUseHandler(itemtypes.ItemTypeWing, itemtypes.ItemWingSubTypeTrialCard, inventoryguaji.GuaJiItemAutoUseHandlerFunc(handleWingAutoUseTrialCard))
}

func handleWingAutoUseTrialCard(pl player.Player, index int32, num int32) {
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
	wingManager := pl.GetPlayerDataManager(playertypes.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	wingInfo := wingManager.GetWingInfo()
	wingAdvancedId := wingInfo.AdvanceId
	if wingAdvancedId != 0 {
		return
	}

	inventorylogic.UseItemIndex(pl, inventorytypes.BagTypePrim, index, 1, nil, "")

	return
}
