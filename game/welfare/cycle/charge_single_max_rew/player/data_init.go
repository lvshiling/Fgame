package player

import (
	cyclechargesinglemaxrewtypes "fgame/fgame/game/welfare/cycle/charge_single_max_rew/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 每日充值
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeMaxRew, playerwelfare.ActivityObjInfoInitFunc(cycleSingleChargeMaxRewInitInfo))
}

func cycleSingleChargeMaxRewInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*cyclechargesinglemaxrewtypes.CycleSingleChargeMaxRewInfo)
	info.CycleDay = welfarelogic.CountCycleDay(groupId)
	info.MaxSingleChargeNum = 0
	info.ReceiveRewRecord = []int32{}
	info.CanRewRecord = []int32{}
}
