package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 消费领奖(每多少领奖)
type CostRewInfo struct {
	GoldNum        int64 `json:"goldNum"`        //消费元宝
	LeftConvertNum int64 `json:"leftConvertNum"` //剩余可兑换元宝数
	ReceiveTimes   int32 `json:"receiveTimes"`   //领取次数
}

func (info *CostRewInfo) CountLeftTimes(convertRate int32) int32 {
	// 计算兑换次数
	addTimes := int32(info.LeftConvertNum) / convertRate
	info.LeftConvertNum -= int64(addTimes * convertRate)
	info.ReceiveTimes += addTimes

	return addTimes
}
func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeRewards, welfaretypes.OpenActivityRewardsSubTypeCost, (*CostRewInfo)(nil))
}
