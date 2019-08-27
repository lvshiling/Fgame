package player

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	shopdiscounttypes "fgame/fgame/game/welfare/shopdiscount/types"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 商城促销
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeShopDiscount, welfaretypes.OpenActivityDefaultSubTypeDefault, playerwelfare.ActivityObjInfoInitFunc(shopDiscountInitInfo))
}

func shopDiscountInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*shopdiscounttypes.ShopDiscountInfo)
	info.PeriodChargeNum = 0
}
