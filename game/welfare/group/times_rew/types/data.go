package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 次数奖励
type TimesRewInfo struct {
	Times     int32             `json:"times"`     //参与次数
	RewRecord map[int32][]int32 `json:"rewRecord"` //奖励领取记录
	IsEmail   bool              `json:"isEmail"`   //是否邮件
}

func (info *TimesRewInfo) AddRecord(rewTimes, vip int32) {
	info.RewRecord[vip] = append(info.RewRecord[vip], rewTimes)
}

func (info *TimesRewInfo) IsCanReceiveRewards(rewTimes, vip int32) bool {
	if info.Times < rewTimes {
		return false
	}
	//领取记录
	recordList := info.RewRecord[vip]
	for _, value := range recordList {
		if value == rewTimes {
			return false
		}
	}

	return true
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeGroup, welfaretypes.OpenActivityGroupSubTypeTimesRew, (*TimesRewInfo)(nil))
}
