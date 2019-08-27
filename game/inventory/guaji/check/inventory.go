package check

import (
	"fgame/fgame/game/guaji/guaji"
	playerguaji "fgame/fgame/game/guaji/player"
	guajitypes "fgame/fgame/game/guaji/types"
	inventoryguaji "fgame/fgame/game/inventory/guaji/guaji"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

func init() {
	guaji.RegisterGuaJiCheckHandler(guajitypes.GuaJiCheckTypeBag, guaji.GuaJiCheckHandlerFunc(inventoryGuaJiCheck))
}

func inventoryGuaJiCheck(pl player.Player) {
	//检查是否购买槽位
	guaJiCheckSlots(pl)
	//主背包检查
	guaJiPrimCheck(pl)
	//宝石背包检查
	guaJiGemCheck(pl)

}

func guaJiCheckSlots(pl player.Player) {
	guaJiManager := pl.GetPlayerDataManager(types.PlayerGuaJiManagerType).(*playerguaji.PlayerGuaJiManager)
	autoBuySlotLevel := guaJiManager.GetGlobalValue(guajitypes.GuaJiAutoBuyBagLevel)
	if pl.GetLevel() < autoBuySlotLevel {
		return
	}
	remainNumOfSlots := inventorylogic.GetNumOfRemainBuySlots(pl)
	if remainNumOfSlots > 0 {
		inventorylogic.HandleBuySlots(pl, remainNumOfSlots)
	}
}

func guaJiPrimCheck(pl player.Player) {
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	items := manager.GetBagAll(inventorytypes.BagTypePrim)

	for _, itemObj := range items {
		if itemObj.IsEmpty() {
			continue
		}
		itemTemplate := item.GetItemService().GetItem(int(itemObj.ItemId))
		if itemTemplate == nil {
			continue
		}

		useHandler := inventoryguaji.GetGuaJiItemUseHandler(itemTemplate.GetItemType())
		if useHandler != nil {
			useHandler.UseItem(pl, itemObj.Index, itemObj.Num)
			continue
		}

		//可以背包使用
		if itemTemplate.UseFlag&itemtypes.ItemUseFlagUse != 0 {
			//TODO 客户端检查
			autoUserHandler := inventoryguaji.GetGuaJiItemAutoUseHandler(itemTemplate.GetItemType(), itemTemplate.GetItemSubType())
			if autoUserHandler == nil {
				continue
			}
			autoUserHandler.AutoUseItem(pl, itemObj.Index, itemObj.Num)
			continue
		}

	}
}

func guaJiGemCheck(pl player.Player) {
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	items := manager.GetBagAll(inventorytypes.BagTypeGem)

	for _, itemObj := range items {
		if itemObj.IsEmpty() {
			continue
		}
		itemTemplate := item.GetItemService().GetItem(int(itemObj.ItemId))
		if itemTemplate == nil {
			continue
		}
		useHandler := inventoryguaji.GetGuaJiItemUseHandler(itemTemplate.GetItemType())
		if useHandler == nil {
			continue
		}
		useHandler.UseItem(pl, itemObj.Index, itemObj.Num)
	}
}
