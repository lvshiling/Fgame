package player

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	rewardscosttypes "fgame/fgame/game/welfare/rewards/cost/types"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 每消费奖励
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeRewards, welfaretypes.OpenActivityRewardsSubTypeCost, playerwelfare.ActivityObjInfoInitFunc(RewardsCostInitInfo))
}

func RewardsCostInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*rewardscosttypes.CostRewInfo)
	info.GoldNum = 0
	info.LeftConvertNum = 0
	info.ReceiveTimes = 0
}
