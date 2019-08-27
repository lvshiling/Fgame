package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 每日充值-单笔充值
type CycleSingleChargeInfo struct {
	CycleDay           int32   `json:"cycleDay"`           //当前充值日
	MaxSingleChargeNum int32   `json:"maxSingleChargeNum"` //单笔最大数
	RewRecord          []int32 `json:"rewRecord"`          //领取记录
	IsEmail            bool    `json:"isEmail"`            //是否奖励发放
}

func (info *CycleSingleChargeInfo) IsCanReceiveRewards(needGoldNum int32) bool {
	//条件
	if info.MaxSingleChargeNum < needGoldNum {
		return false
	}

	//领取记录
	for _, value := range info.RewRecord {
		if value == needGoldNum {
			return false
		}
	}

	return true
}

func (info *CycleSingleChargeInfo) AddRecord(needGoldNum int32) {
	info.RewRecord = append(info.RewRecord, needGoldNum)
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleCharge, (*CycleSingleChargeInfo)(nil))
}
