package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//商城促销
type ShopDiscountInfo struct {
	PeriodChargeNum int32 `json:"periodChargeNum"` //活动期间充值数量
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeShopDiscount, welfaretypes.OpenActivityDefaultSubTypeDefault, (*ShopDiscountInfo)(nil))
}
