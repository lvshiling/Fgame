package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playershop "fgame/fgame/game/shop/player"
	"fgame/fgame/game/shop/shop"
)

func BuildSCShopLimit(buyCountMap map[int32]*playershop.PlayerShopObject) *uipb.SCShopLimit {
	scShopLimit := &uipb.SCShopLimit{}
	for _, shop := range buyCountMap {
		scShopLimit.ShopLimitList = append(scShopLimit.ShopLimitList, buildShop(shop))
	}
	return scShopLimit
}

func BuildSCShopBuy(shopId int32, num int32, dayCount int32) *uipb.SCShopBuy {
	shopBuy := &uipb.SCShopBuy{}
	shopBuy.ShopId = &shopId
	shopBuy.Num = &num
	shopBuy.DayCount = &dayCount
	return shopBuy
}

func BuildSCShopAutoBuyList(pl player.Player, shopIdMap map[int32]int32) *uipb.SCShopAutoBuyList {
	scShopAutoBuyList := &uipb.SCShopAutoBuyList{}
	manager := pl.GetPlayerDataManager(types.PlayerShopDataManagerType).(*playershop.PlayerShopDataManager)
	for shopId, num := range shopIdMap {
		shopTemplate := shop.GetShopService().GetShopTemplate(int(shopId))
		if shopTemplate == nil {
			continue
		}
		dayCount := manager.GetDayCountByShopId(shopId)
		scShopAutoBuyList.AutoBuyList = append(scShopAutoBuyList.AutoBuyList, buildAutoBuy(shopId, num, dayCount))
	}
	return scShopAutoBuyList
}

func BuildSCShopStopAutoBuy() *uipb.SCShopStopAutoBuy {
	scShopStopAutoBuy := &uipb.SCShopStopAutoBuy{}
	return scShopStopAutoBuy
}

func buildAutoBuy(shopId int32, num int32, dayCount int32) *uipb.ShouAutoBuy {
	shouAutoBuy := &uipb.ShouAutoBuy{}
	shouAutoBuy.ShopId = &shopId
	shouAutoBuy.Num = &num
	shouAutoBuy.DayCount = &dayCount
	return shouAutoBuy
}

func buildShop(shop *playershop.PlayerShopObject) *uipb.ShopLimit {
	shopLimit := &uipb.ShopLimit{}
	shopId := int32(shop.ShopId)
	num := int32(shop.DayCount)
	shopLimit.ShopId = &shopId
	shopLimit.DayCount = &num
	return shopLimit
}
