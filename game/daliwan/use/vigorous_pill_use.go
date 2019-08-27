package use

import (
	playerdailiwan "fgame/fgame/game/daliwan/player"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeVigorousPill, itemtypes.ItemDefaultSubTypeDefault, playerinventory.ItemUseHandleFunc(handleVigorousPill))
}

func handleVigorousPill(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	tempId := itemTemplate.TypeFlag1
	manager := pl.GetPlayerDataManager(playertypes.PlayerDaLiWanDataManagerType).(*playerdailiwan.PlayerDaLiWanManager)
	flag = manager.Use(tempId)
	if !flag {
		return false, nil
	}

	return
}
