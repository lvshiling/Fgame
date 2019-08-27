package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//福利-在线
type WelfareOnlineInfo struct {
	DrawTimes int32   `json:"drawTimes"` //抽奖次数
	RewRecord []int32 `json:"rewRecord"` //奖励领取记录
}

func (info *WelfareOnlineInfo) AddRecord(rewTime int32) {
	info.RewRecord = append(info.RewRecord, rewTime)
}

func (info *WelfareOnlineInfo) IsReceive(rewTime int32) bool {
	for _, value := range info.RewRecord {
		if value == rewTime {
			return true
		}
	}
	return false
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeOnline, (*WelfareOnlineInfo)(nil))
}
