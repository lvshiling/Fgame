package use

import (
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/unrealboss/pbutil"
	playerunrealboss "fgame/fgame/game/unrealboss/player"
)

func init() {
	playerinventory.RegisterUseHandler(itemtypes.ItemTypeBossItem, itemtypes.ItemBossSubTypeXingShenDan, playerinventory.ItemUseHandleFunc(handleUseBossItemXingShengDan))
}

// 使用醒神丹
func handleUseBossItemXingShengDan(pl player.Player, it *playerinventory.PlayerItemObject, itemNum int32, chooseIndexList []int32, args string) (flag bool, err error) {
	itemId := it.ItemId
	itemTemplate := item.GetItemService().GetItem(int(itemId))
	pilao := int32(itemTemplate.TypeFlag1)

	// 设置疲劳值
	unrealBossManager := pl.GetPlayerDataManager(playertypes.PlayerUnrealBossDataManagerType).(*playerunrealboss.PlayerUnrealBossDataManager)
	unrealBossManager.SetPiLaoByXingShenDan(pilao)

	scMsg := pbutil.BuildSCUnrealBossBuyPilaoNum(unrealBossManager.GetCurPilaoNum())
	pl.SendMsg(scMsg)

	flag = true
	return
}
