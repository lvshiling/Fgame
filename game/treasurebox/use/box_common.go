package use

import (
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	boxlogic "fgame/fgame/game/treasurebox/logic"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeGiftBag, itemtypes.ItemGiftBagSubTypeBox, playerinventory.ItemUseHandleFunc(handleGiftBagUse))
}

func handleGiftBagUse(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	startBoxTemplate := itemTemplate.GetBoxTemplate()

	return boxlogic.OpenBox(pl, itemId, num, chooseIndexList, startBoxTemplate)
}
