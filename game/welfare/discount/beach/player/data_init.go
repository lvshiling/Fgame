package player

import (
	discountbeachtypes "fgame/fgame/game/welfare/discount/beach/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeBeach, playerwelfare.ActivityObjInfoInitFunc(discountBeachShopInitInfo))
}

func discountBeachShopInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*discountbeachtypes.BeachShopInfo)
	info.BuyRecord = make(map[int32]int32)
	info.IsActivite = 0
}
