package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//福利-升级
type WelfareUplevelInfo struct {
	RewRecord []int32 `json:"rewRecord"` //奖励领取记录
	IsEmail   bool    `json:"isEmail"`   //是否奖励发放
}

func (info *WelfareUplevelInfo) AddRecord(rewLevel int32) {
	info.RewRecord = append(info.RewRecord, rewLevel)
}

func (info *WelfareUplevelInfo) IsReceive(rewLevel int32) bool {
	for _, value := range info.RewRecord {
		if value == rewLevel {
			return true
		}
	}
	return false
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeUpLevel, (*WelfareUplevelInfo)(nil))
}
