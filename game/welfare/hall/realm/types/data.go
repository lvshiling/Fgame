package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//福利-天劫塔冲刺
type WelfareRealmChallengeInfo struct {
	Level     int32   `json:"level"`     //天劫塔等级
	RewRecord []int32 `json:"rewRecord"` //领取记录
	IsEmail   bool    `json:"isEmail"`   //是否奖励发放
}

func (info *WelfareRealmChallengeInfo) IsCanReceiveRewards(needLevel int32) bool {
	//条件
	if info.Level < needLevel {
		return false
	}

	//领取记录
	for _, value := range info.RewRecord {
		if value == needLevel {
			return false
		}
	}

	return true
}

func (info *WelfareRealmChallengeInfo) AddRecord(needLevel int32) {
	info.RewRecord = append(info.RewRecord, needLevel)
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeRealm, (*WelfareRealmChallengeInfo)(nil))
}
