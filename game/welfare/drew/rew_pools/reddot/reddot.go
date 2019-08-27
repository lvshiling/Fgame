package reddot

import (
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/reddot/reddot"
	rewpoolstypes "fgame/fgame/game/welfare/drew/rew_pools/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeRewPools, reddot.HandlerFunc(handleRedDotRewPools))
}

func handleRedDotRewPools(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	welfareTemplateService := welfaretemplate.GetWelfareTemplateService()
	groupInterface := welfareTemplateService.GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}

	info := obj.GetActivityData().(*rewpoolstypes.RewPoolsInfo)
	luckDrewTemp := welfareTemplateService.GetLuckDrewTemplateByArg(groupId, info.Position)
	if luckDrewTemp == nil {
		return
	}
	for itemId, itemNeedNum := range luckDrewTemp.GetUseItemMap() {
		if inventoryManager.NumOfItems(itemId) < itemNeedNum {
			return false
		}
	}
	return true
}
