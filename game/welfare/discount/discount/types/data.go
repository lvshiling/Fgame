package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//限时折扣
type DiscountInfo struct {
	BuyRecord   map[int32]int32 `json:"buyRecord"`   //购买记录
	DiscountDay int32           `json:"discountDay"` //折扣日
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeCommon, (*DiscountInfo)(nil))
}
