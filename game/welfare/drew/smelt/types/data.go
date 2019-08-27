package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

type SmeltInfo struct {
	Num           int32 `json:"num"`              //数量
	HasReceiveNum int32 `json:"hasReceiveRecord"` //已经领取的记录次数
}

func (info *SmeltInfo) GetRemainCanReceiveRecord(needItemNum int32) int32 {
	if needItemNum == 0 {
		return 0
	}
	canReceiveNum := info.Num / needItemNum
	remain := canReceiveNum - info.HasReceiveNum
	if remain < 0 {
		remain = 0
	}
	return remain
}

func (info *SmeltInfo) IsCanReceiveReward(needItemNum int32) bool {
	if needItemNum == 0 {
		return false
	}
	canReceiveNum := info.Num / needItemNum
	remain := canReceiveNum - info.HasReceiveNum
	if remain <= 0 {
		return false
	}
	return true
}

func (info *SmeltInfo) AddReceiveRecord(times int32) {
	info.HasReceiveNum = info.HasReceiveNum + times
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeSmelt, (*SmeltInfo)(nil))
}
