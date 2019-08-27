package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

const (
	defaultCycleNum = 1
)

//充值抽奖
type LuckyChargeDrewInfo struct {
	GoldNum         int32 `json:"goldNum"`         //充值元宝
	LeftConvertNum  int32 `json:"leftConvertNum"`  //剩余可兑换元宝数
	HadConvertTimes int32 `json:"hadConvertTimes"` //已兑换次数
	LeftTimes       int32 `json:"leftTimes"`       //剩余可参与次数
	AttendTimes     int32 `json:"attendTimes"`     //已参与次数
	Ratio           int32 `json:"ratio"`           //奖励倍数
	CycleCount      int32 `json:"cycleCount"`      //后台规则计数
}

func (info *LuckyChargeDrewInfo) CountLeftTimes(convertLimit, convertRate, minCycle int32) {
	if convertRate == 0 {
		return
	}

	if minCycle == 0 {
		minCycle = defaultCycleNum
	}

	// 计算兑换次数
	if info.HadConvertTimes < convertLimit {
		timesRaio := info.LeftConvertNum / (convertRate * minCycle)
		addTimes := timesRaio * minCycle
		curMaxTimes := convertLimit - info.HadConvertTimes
		if addTimes > curMaxTimes {
			addTimes = curMaxTimes
		}

		info.LeftConvertNum -= addTimes * convertRate
		info.LeftTimes += addTimes
		info.HadConvertTimes += addTimes
	}
}

func (info *LuckyChargeDrewInfo) ResetRule() {
	info.Ratio = 0
	info.CycleCount = 0
}

func (info *LuckyChargeDrewInfo) IsRuleCD(ruleCycle int32) bool {
	return info.CycleCount <= ruleCycle
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeChargeDrew, (*LuckyChargeDrewInfo)(nil))
}
