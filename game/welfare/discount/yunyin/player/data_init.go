package player

import (
	discountyunyintypes "fgame/fgame/game/welfare/discount/yunyin/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeDiscount, welfaretypes.OpenActivityDiscountSubTypeYunYin, playerwelfare.ActivityObjInfoInitFunc(yunYinInitInfo))
}

func yunYinInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*discountyunyintypes.YunYinInfo)
	info.BuyRecord = make(map[int32]int32)
	info.ReceiveRecord = make([]int32, 0, 8)
	info.GoldNum = 0
	info.IsEmail = false
}
