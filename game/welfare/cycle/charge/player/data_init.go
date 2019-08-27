package player

import (
	cyclechargetypes "fgame/fgame/game/welfare/cycle/charge/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 每日充值
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeCharge, playerwelfare.ActivityObjInfoInitFunc(cycleChargeInitInfo))
}

func cycleChargeInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*cyclechargetypes.CycleChargeInfo)
	info.CycleDay = welfarelogic.CountCycleDay(groupId)
	info.GoldNum = 0
	info.RewRecord = []int32{}
}
