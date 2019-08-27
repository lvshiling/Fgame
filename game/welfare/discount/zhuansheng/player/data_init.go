package player

import (
	discountzhuanshengtypes "fgame/fgame/game/welfare/discount/zhuansheng/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 限时折扣
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeZhuanSheng, playerwelfare.ActivityObjInfoInitFunc(discountZhuanShengInitInfo))
}

func discountZhuanShengInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*discountzhuanshengtypes.DiscountZhuanShengInfo)
	info.BuyRecord = make(map[int32]int32)
	info.GiftReceiveRecord = make(map[int32]int32)
	info.ChargeNum = 0
	info.UsePoint = 0
}
