package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 每日充值-单笔充值（只领取最高奖励）
type CycleSingleChargeMaxRewInfo struct {
	CycleDay           int32   `json:"cycleDay"`           //当前充值日
	MaxSingleChargeNum int32   `json:"maxSingleChargeNum"` //单笔最大数
	CanRewRecord       []int32 `json:"canRewRecord"`       //可领取记录
	ReceiveRewRecord   []int32 `json:"receiveRewRecord"`   //已领取记录
}

func (info *CycleSingleChargeMaxRewInfo) IsCanReceiveRewards(needGoldNum int32) bool {
	//条件
	flag := false
	for _, value := range info.CanRewRecord {
		if value != needGoldNum {
			continue
		}
		flag = true
	}

	return flag
}

func (info *CycleSingleChargeMaxRewInfo) AddReceiveRecord(needGoldNum int32) {

	//删除 可领取记录
	delIndex := -1
	for index, value := range info.CanRewRecord {
		if value == needGoldNum {
			delIndex = index
			break
		}
	}
	if delIndex == -1 {
		return
	}
	info.CanRewRecord = append(info.CanRewRecord[:delIndex], info.CanRewRecord[delIndex+1:]...)

	// 添加 已领取记录
	info.ReceiveRewRecord = append(info.ReceiveRewRecord, needGoldNum)
}

func (info *CycleSingleChargeMaxRewInfo) AddCanRewRecord(needGoldNum int32) {
	info.CanRewRecord = append(info.CanRewRecord, needGoldNum)
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeMaxRew, (*CycleSingleChargeMaxRewInfo)(nil))
}
