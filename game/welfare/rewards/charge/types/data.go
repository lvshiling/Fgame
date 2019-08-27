package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 充值领奖(每多少领奖)
type ChargeRewInfo struct {
	GoldNum        int32 `json:"goldNum"`        //充值元宝
	LeftConvertNum int32 `json:"leftConvertNum"` //剩余可兑换元宝数(废弃)
	ReceiveTimes   int32 `json:"receiveTimes"`   //领取次数
	Fixed          bool  `json:"fixed"`
}

func (info *ChargeRewInfo) CountLeftTimes(convertRate int32) int32 {
	// 计算兑换次数
	leftGold := info.GoldNum - info.ReceiveTimes*convertRate
	addTimes := leftGold / convertRate
	info.ReceiveTimes += addTimes

	return addTimes
}

func (info *ChargeRewInfo) FixMore(goldNum int32, convertRate int32) {
	// 计算兑换次数
	info.ReceiveTimes -= goldNum / convertRate
	if info.ReceiveTimes <= 0 {
		info.ReceiveTimes = 0
	}
	info.Fixed = true

}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeRewards, welfaretypes.OpenActivityRewardsSubTypeCharge, (*ChargeRewInfo)(nil))
}
