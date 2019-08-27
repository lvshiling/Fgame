package player

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 限时砍价
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeTaoCan, playerwelfare.ActivityObjInfoInitFunc(discountTaoCanInitInfo))
}

func discountTaoCanInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
}
