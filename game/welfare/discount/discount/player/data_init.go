package player

import (
	discountdiscounttypes "fgame/fgame/game/welfare/discount/discount/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 限时折扣
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeCommon, playerwelfare.ActivityObjInfoInitFunc(discountInitInfo))
}

func discountInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*discountdiscounttypes.DiscountInfo)
	info.BuyRecord = make(map[int32]int32)
	info.DiscountDay = welfarelogic.CountCurActivityDay(groupId)
}
