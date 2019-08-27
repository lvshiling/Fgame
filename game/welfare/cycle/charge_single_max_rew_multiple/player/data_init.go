package player

import (
	cyclechargesinglemaxrewmultipletypes "fgame/fgame/game/welfare/cycle/charge_single_max_rew_multiple/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 每日充值
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeMaxRewMultiple, playerwelfare.ActivityObjInfoInitFunc(cycleSingleChargeMaxRewMultipleInitInfo))
}

func cycleSingleChargeMaxRewMultipleInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*cyclechargesinglemaxrewmultipletypes.CycleSingleChargeMaxRewMultipleInfo)
	info.CycleDay = welfarelogic.CountCycleDay(groupId)
	info.MaxSingleChargeNum = 0
	info.ReceiveRewRecord = []int32{}
	info.CanRewRecord = []int32{}
	info.NewReceiveRewRecord = make(map[int32]int32)
	info.NewCanRewRecord = make(map[int32]int32)
}
