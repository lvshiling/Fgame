package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//累充奖励
type FeedbackChargeReturnMultipleInfo struct {
	PeriodChargeNum int32 `json:"periodChargeNum"` //活动期间充值数量
	RewardCnt       int32 `json:"rewardCnt"`       //已领取次数
}

//增加活动期间充值数量
func (info *FeedbackChargeReturnMultipleInfo) AddPeriodCharge(addGold int32) {
	info.PeriodChargeNum += addGold
}

//增加活动期间充值奖励领取次数
func (info *FeedbackChargeReturnMultipleInfo) AddRewardCnt(val int32) {
	info.RewardCnt += val
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeReturnMultiple, (*FeedbackChargeReturnMultipleInfo)(nil))
}
