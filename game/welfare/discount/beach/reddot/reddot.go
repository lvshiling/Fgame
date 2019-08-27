package reddot

import (
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/reddot/reddot"
	discountbeachtemplate "fgame/fgame/game/welfare/discount/beach/template"
	discountbeachtypes "fgame/fgame/game/welfare/discount/beach/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeBeach, reddot.HandlerFunc(handleReddotBeachShop))
}

func handleReddotBeachShop(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}

	// 判断商店是否激活
	info := obj.GetActivityData().(*discountbeachtypes.BeachShopInfo)
	if info.IsActivited() {
		return
	}

	// 判断是否有足够的物品
	beachTemp := groupInterface.(*discountbeachtemplate.GroupTemplateDiscountBeachShop)
	itemMap := beachTemp.GetAvtiviteItemMap()
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if len(itemMap) != 0 {
		if !inventoryManager.HasEnoughItems(itemMap) {
			return
		}
	}
	isNotice = true
	return
}
