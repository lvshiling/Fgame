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
)

func init() {
	inventoryguaji.RegisterGuaJiItemUseHandler(itemtypes.ItemTypeFuChi, inventoryguaji.GuaJiItemUseHandlerFunc(handleFuChiUse))
}

func handleFuChiUse(pl player.Player, index int32, num int32) {
	//参数不对
	if num != 1 {
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
	if itemTemplate == nil {
		return
	}

	//使用装备
	inventorylogic.UseItemIndex(pl, inventorytypes.BagTypePrim, index, num, nil, "")
}
