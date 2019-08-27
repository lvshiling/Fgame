package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 每日首充
type CycleChargeInfo struct {
	GoldNum   int32   `json:"goldNum"`   //充值元宝数量
	RewRecord []int32 `json:"rewRecord"` //奖励领取记录
	CycleDay  int32   `json:"cycleDay"`  //当前充值日
}

func (info *CycleChargeInfo) AddRecord(needGoldNum int32) {
	info.RewRecord = append(info.RewRecord, needGoldNum)
}

func (info *CycleChargeInfo) IsCanReceiveRewards(needGoldNum int32) bool {
	//条件
	if info.GoldNum < needGoldNum {
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

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeCharge, (*CycleChargeInfo)(nil))
}
