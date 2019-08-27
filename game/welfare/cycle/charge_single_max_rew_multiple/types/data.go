package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 每日充值-单笔充值（只领取最高奖励,多次）
type CycleSingleChargeMaxRewMultipleInfo struct {
	CycleDay            int32           `json:"cycleDay"`            //当前充值日
	MaxSingleChargeNum  int32           `json:"maxSingleChargeNum"`  //单笔最大数
	CanRewRecord        []int32         `json:"canRewRecord"`        //可领取记录 //废弃
	ReceiveRewRecord    []int32         `json:"receiveRewRecord"`    //已领取记录 //废弃
	NewCanRewRecord     map[int32]int32 `json:"newCanRewRecord"`     //新可领取记录
	NewReceiveRewRecord map[int32]int32 `json:"newReceiveRewRecord"` //新已领取记录
}

func (info *CycleSingleChargeMaxRewMultipleInfo) IsCanReceiveRewards(needGoldNum int32) bool {
	//条件
	leftNum, ok := info.NewCanRewRecord[needGoldNum]
	if !ok {
		return false
	}

	getNum := info.NewReceiveRewRecord[needGoldNum]
	if leftNum <= getNum {
		return false
	}

	return true
}

func (info *CycleSingleChargeMaxRewMultipleInfo) LeftCanReceiveRewards() map[int32]int32 {
	leftM := map[int32]int32{}
	for key, val := range info.NewCanRewRecord {
		getNum := info.NewReceiveRewRecord[key]
		leftNum := val - getNum
		if leftNum > 0 {
			leftM[key] = leftNum
		}
	}
	return leftM
}

func (info *CycleSingleChargeMaxRewMultipleInfo) AddReceiveRecord(needGoldNum int32) {

	//删除 可领取记录
	// leftNum, ok := info.NewCanRewRecord[needGoldNum]
	// if !ok {
	// 	return
	// }
	// leftNum -= 1

	// if leftNum <= 0 {
	// 	delete(info.NewCanRewRecord, needGoldNum)
	// } else {
	// 	info.NewCanRewRecord[needGoldNum] = leftNum
	// }

	// 添加 已领取记录
	getNum := info.NewReceiveRewRecord[needGoldNum]
	getNum += 1
	info.NewReceiveRewRecord[needGoldNum] = getNum
}

func (info *CycleSingleChargeMaxRewMultipleInfo) AddCanRewRecord(needGoldNum int32) {
	leftNum := info.NewCanRewRecord[needGoldNum]
	leftNum += 1
	info.NewCanRewRecord[needGoldNum] = leftNum
}

//改新字段了
func (info *CycleSingleChargeMaxRewMultipleInfo) ChangeNewFields() {
	if info.NewCanRewRecord != nil {
		return
	}

	info.NewReceiveRewRecord = make(map[int32]int32)
	info.NewCanRewRecord = make(map[int32]int32)
	for _, val := range info.ReceiveRewRecord {
		_, ok := info.NewReceiveRewRecord[val]
		if ok {
			info.NewReceiveRewRecord[val] += 1
		} else {
			info.NewReceiveRewRecord[val] = 1
		}
		_, flag := info.NewCanRewRecord[val]
		if flag {
			info.NewCanRewRecord[val] += 1
		} else {
			info.NewCanRewRecord[val] = 1
		}
	}
	for _, val := range info.CanRewRecord {
		_, ok := info.NewCanRewRecord[val]
		if ok {
			info.NewCanRewRecord[val] += 1
		} else {
			info.NewCanRewRecord[val] = 1
		}
	}

	// info.ReceiveRewRecord = []int32{}
	// info.CanRewRecord = []int32{}
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeMaxRewMultiple, (*CycleSingleChargeMaxRewMultipleInfo)(nil))
}
