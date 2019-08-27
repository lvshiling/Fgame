package player

import (
	drewchargedrewtypes "fgame/fgame/game/welfare/drew/charge_drew/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 充值抽奖
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeChargeDrew, playerwelfare.ActivityObjInfoInitFunc(chargeDrewInitInfo))
}

func chargeDrewInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*drewchargedrewtypes.LuckyChargeDrewInfo)
	info.AttendTimes = 0
	info.GoldNum = 0
	info.LeftConvertNum = 0
	info.HadConvertTimes = 0
	info.LeftTimes = 0
	info.Ratio = 0
}
