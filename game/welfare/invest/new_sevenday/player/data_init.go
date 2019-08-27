package player

import (
	investnewsevendaytypes "fgame/fgame/game/welfare/invest/new_sevenday/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 新七日投资
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeInvest, welfaretypes.OpenActivityInvestSubTypeNewServenDay, playerwelfare.ActivityObjInfoInitFunc(investDayInitInfo))
}

func investDayInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*investnewsevendaytypes.NewInvestDayInfo)
	info.ReceiveMap = make(map[int32]int32)
	info.BuyTimeMap = make(map[int32]int64)
	info.MaxSingleChargeNum = 0
	info.IsEmail = false
}
