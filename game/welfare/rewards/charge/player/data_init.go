package player

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	rewardschargetypes "fgame/fgame/game/welfare/rewards/charge/types"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 每充值奖励
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeRewards, welfaretypes.OpenActivityRewardsSubTypeCharge, playerwelfare.ActivityObjInfoInitFunc(RewardsChargeInitInfo))
}

func RewardsChargeInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*rewardschargetypes.ChargeRewInfo)
	info.GoldNum = 0
	info.LeftConvertNum = 0
	info.ReceiveTimes = 0
	info.Fixed = true
}
