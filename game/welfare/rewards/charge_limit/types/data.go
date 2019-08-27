package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fmt"
)

// 充值领奖(每多少领奖)（全服次数）
type ChargeRewLimitInfo struct {
	GoldNum           int32           `json:"goldNum"`           //充值元宝
	LeftConvertNumMap map[int32]int32 `json:"leftConvertNumMap"` //剩余可兑换元宝数（每个档次）
	TotalReceiveTimes int32           `json:"receiveTimes"`      //总领取次数
	ReceiveTimesMap   map[int32]int32 `json:"receiveTimesMap"`   //领取次数（每个档次）
}

// ReceiveTimesMap 新增字段，要初始化
func (info *ChargeRewLimitInfo) Fix() bool {
	if info.ReceiveTimesMap == nil {
		info.ReceiveTimesMap = make(map[int32]int32)
		return true
	}

	return false
}

func (info *ChargeRewLimitInfo) CountLeftTimes(convertRate, playerMaxTimes int32) int32 {
	if convertRate == 0 {
		panic(fmt.Errorf("兑换数不能为0，convertRate:%d", convertRate))
	}

	// 计算兑换次数
	leftGoldNum, ok := info.LeftConvertNumMap[convertRate]
	if !ok {
		return 0
	}

	playerTimes := info.ReceiveTimesMap[convertRate]

	rewNum := leftGoldNum / convertRate
	if rewNum+playerTimes <= playerMaxTimes {
		return rewNum
	} else {
		return playerMaxTimes - playerTimes
	}
}

func (info *ChargeRewLimitInfo) AddLeftNum(goldNum int32) {
	for convertRate, _ := range info.LeftConvertNumMap {
		info.LeftConvertNumMap[convertRate] += goldNum
	}
}

func (info *ChargeRewLimitInfo) ReceiveRewards(convertRate, addTimes int32) {
	info.LeftConvertNumMap[convertRate] -= addTimes * convertRate
	info.ReceiveTimesMap[convertRate] += addTimes
	info.TotalReceiveTimes += addTimes
}

func (info *ChargeRewLimitInfo) IsHadReceiveTimes(convertRate, timesMax, rewTimes int32) bool {
	times := info.ReceiveTimesMap[convertRate]
	if times+rewTimes > timesMax {
		return false
	}

	return true
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeRewards, welfaretypes.OpenActivityRewardsSubTypeChargeLimit, (*ChargeRewLimitInfo)(nil))
}
