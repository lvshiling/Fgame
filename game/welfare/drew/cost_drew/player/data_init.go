package player

import (
	drewcostdrewtypes "fgame/fgame/game/welfare/drew/cost_drew/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 消费抽奖
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeCostDrew, playerwelfare.ActivityObjInfoInitFunc(chargeDrewInitInfo))
}

func chargeDrewInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*drewcostdrewtypes.LuckyCostDrewInfo)
	info.AttendTimes = 0
	info.GoldNum = 0
	info.LeftConvertNum = 0
	info.HadConvertTimes = 0
	info.LeftTimes = 0
	info.Ratio = 0
}
