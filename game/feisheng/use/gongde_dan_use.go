package use

import (
	"fgame/fgame/game/feisheng/pbutil"
	playerfeisheng "fgame/fgame/game/feisheng/player"
	funcopentypes "fgame/fgame/game/funcopen/types"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeFeiSheng, itemtypes.ItemFeiShengSubTypeGongDeDan, playerinventory.ItemUseHandleFunc(handleUseDan))
}

//功德丹
func handleUseDan(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeFeiSheng) {
		return
	}

	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	addGongDe := int64(itemTemplate.TypeFlag1 * num)
	feiManager := pl.GetPlayerDataManager(playertypes.PlayerFeiShengDataManagerType).(*playerfeisheng.PlayerFeiShengDataManager)
	feiManager.AddGongDe(addGongDe)

	feiShengInfo := feiManager.GetFeiShengInfo()
	scMsg := pbutil.BuildSCFeiShengInfo(feiShengInfo)
	pl.SendMsg(scMsg)

	flag = true
	return
}
