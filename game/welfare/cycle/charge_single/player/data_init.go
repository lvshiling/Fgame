package player

import (
	cyclechargesingletypes "fgame/fgame/game/welfare/cycle/charge_single/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 每日充值
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleCharge, playerwelfare.ActivityObjInfoInitFunc(cycleSingleChargeInitInfo))
}

func cycleSingleChargeInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*cyclechargesingletypes.CycleSingleChargeInfo)
	info.CycleDay = welfarelogic.CountCycleDay(groupId)
	info.MaxSingleChargeNum = 0
	info.RewRecord = []int32{}
	info.IsEmail = false
}
