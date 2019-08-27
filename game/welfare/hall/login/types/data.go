package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//福利-登录
type WelfareLoginInfo struct {
	RewRecord []int32 `json:"rewRecord"` //奖励领取记录
}

func (info *WelfareLoginInfo) AddRecord(rewDay int32) {
	info.RewRecord = append(info.RewRecord, rewDay)
}

func (info *WelfareLoginInfo) IsReceive(rewDay int32) bool {
	for _, value := range info.RewRecord {
		if value == rewDay {
			return true
		}
	}
	return false
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeLogin, (*WelfareLoginInfo)(nil))
}
