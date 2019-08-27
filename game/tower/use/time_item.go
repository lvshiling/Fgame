package use

import (
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/tower/pbutil"
	playertower "fgame/fgame/game/tower/player"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeTower, itemtypes.ItemTowerSubTypeTimeCard, playerinventory.ItemUseHandleFunc(handleAdddTowerTime))
}

//时间沙漏
func handleAdddTowerTime(pl player.Player, it *playerinventory.PlayerItemObject, num int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemp := item.GetItemService().GetItem(int(itemId))
	extraTime := int64(itemTemp.TypeFlag1) * int64(num)

	towerManager := pl.GetPlayerDataManager(playertypes.PlayerTowerDataManagerType).(*playertower.PlayerTowerDataManager)
	towerManager.AddExtraTime(extraTime)
	remainTime := towerManager.GetRemainTime()

	scMsg := pbutil.BuildSCTowerTimeNotice(remainTime)
	pl.SendMsg(scMsg)

	flag = true
	return
}
