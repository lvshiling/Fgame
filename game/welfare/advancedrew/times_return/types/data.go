package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//升阶次数返还
type AdvancedTimesReturnInfo struct {
	Times     int32                     `json:"times"`        //进阶次数
	RewRecord []int32                   `json:"rewRecord"`    //奖励领取记录
	RewType   welfaretypes.AdvancedType `json:"advancedType"` //系统类型
	IsEmail   bool                      `json:"isEmail"`      //是否邮件
}

func (info *AdvancedTimesReturnInfo) AddRecord(needTimes int32) {
	info.RewRecord = append(info.RewRecord, needTimes)
}

func (info *AdvancedTimesReturnInfo) IsCanReceiveRewards(needTimes int32) bool {
	//条件
	if info.Times < needTimes {
		return false
	}

	//领取记录
	for _, value := range info.RewRecord {
		if value == needTimes {
			return false
		}
	}

	return true
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeTimesReturn, (*AdvancedTimesReturnInfo)(nil))
}
