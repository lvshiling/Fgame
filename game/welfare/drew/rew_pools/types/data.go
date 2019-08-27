package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

type RewPoolsInfo struct {
	Position   int32 `json:"position"`   //奖池位置
	BackTimes  int32 `json:"backTimes"`  //回退次数
	RecordTime int64 `json:"recordTime"` //记录时间
}

func (info *RewPoolsInfo) PoolForWard() {
	info.Position = info.Position + 1
}

func (info *RewPoolsInfo) BackTimesEnoughPoolForWard(loopTimes int32) {
	info.Position = info.Position + 1
	info.BackTimes = info.BackTimes % loopTimes
}

func (info *RewPoolsInfo) PoolBack() {
	info.Position = info.Position - 1
	info.BackTimes = info.BackTimes + 1
}

func (info *RewPoolsInfo) GetBackTimes() int32 {
	return info.BackTimes
}

func (info *RewPoolsInfo) IsBackTimesEnough(loopTimes int32) bool {
	if info.BackTimes%loopTimes == 0 && info.BackTimes > 0 {
		return true
	} else {
		return false
	}
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeRewPools, (*RewPoolsInfo)(nil))
}
