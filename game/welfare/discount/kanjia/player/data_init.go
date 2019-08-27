package player

import (
	discountkanjiatypes "fgame/fgame/game/welfare/discount/kanjia/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 限时砍价
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeKanJia, playerwelfare.ActivityObjInfoInitFunc(discountKanJiaInitInfo))
}

func discountKanJiaInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*discountkanjiatypes.DiscountKanJiaInfo)
	info.UseTimes = 0
	info.GoldNum = 0
	info.BuyRecord = []int32{}
	info.KanJiaRecord = make(map[int32]*discountkanjiatypes.KanJiaInfo)
}
