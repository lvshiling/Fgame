package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

type CycleChargeSingleAllMultipleInfo struct {
	CycleDay           int32           `json:"cycleDay"`           //当前充值日
	SingleChargeRecord []int32         `json:"singleChargeRecord"` //单笔最大数 (废弃)
	CanRewRecord       map[int32]int32 `json:"canRewRecord"`       //可领取记录 (废弃)

	NewSingleChargeRecord map[int32]int32 `json:"newSingleChargeRecord"` //单笔最大数
	RewRecord             map[int32]int32 `json:"rewRecord"`             //已经领取的
}

func (info *CycleChargeSingleAllMultipleInfo) AddSingleChargeRecord(needGoldNum int32) {
	info.NewSingleChargeRecord[needGoldNum] += 1
}

// func (info *CycleChargeSingleAllMultipleInfo) AddCanRewRecord(needGoldNum int32, times int32) {
// 	info.CanRewRecord[needGoldNum] += times
// }

func (info *CycleChargeSingleAllMultipleInfo) GetCanRewRecord() map[int32]int32 {
	remainTimeMap := make(map[int32]int32)
	for goldNum, times := range info.NewSingleChargeRecord {
		useTimes := info.RewRecord[goldNum]
		remainTimes := times - useTimes
		if remainTimes <= 0 {
			continue
		}
		remainTimeMap[goldNum] = remainTimes
	}
	return remainTimeMap
}

func (info *CycleChargeSingleAllMultipleInfo) Receive(useTimes map[int32]int32) {
	for goldNum, times := range useTimes {
		info.RewRecord[goldNum] += times
	}
	// info.SingleChargeRecord = make([]int32, 1)
	// for key, _ := range info.CanRewRecord {
	// 	info.CanRewRecord[key] = 0
	// }
}

func (info *CycleChargeSingleAllMultipleInfo) SyncCharges(chargeMap map[int32]int32) bool {
	changed := false
	for goldNum, timesNum := range chargeMap {
		if info.NewSingleChargeRecord[goldNum] != timesNum {
			info.NewSingleChargeRecord[goldNum] = timesNum
			changed = true
		}
	}
	return changed
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeAllRew, (*CycleChargeSingleAllMultipleInfo)(nil))
}
