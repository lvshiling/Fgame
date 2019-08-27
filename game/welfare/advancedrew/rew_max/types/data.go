package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//进阶奖励（活动开始初始化进阶等级，只能领取当前最近档次及后面的阶数奖励）
type AdvancedRewMaxInfo struct {
	RewType         welfaretypes.AdvancedType `json:"rewType"`         //进阶类型
	AdvancedNum     int32                     `json:"advancedNum"`     //阶数
	RewRecord       []int32                   `json:"rewRecord"`       //奖励领取记录
	IsEmail         bool                      `json:"isEmail"`         //是否邮件
	PeriodChargeNum int32                     `json:"periodChargeNum"` //期间充值数
	InitAdvancedNum int32                     `json:"initAdvancedNum"` //初始阶数
}

func (info *AdvancedRewMaxInfo) AddRecord(needAdvancedNum int32) {
	info.RewRecord = append(info.RewRecord, needAdvancedNum)
}

func (info *AdvancedRewMaxInfo) IsCanReceiveRewards(needAdvancedNum, needChargeNum int32) bool {
	if info.AdvancedNum < needAdvancedNum {
		return false
	}
	if info.PeriodChargeNum < needChargeNum {
		return false
	}
	//领取记录
	for _, value := range info.RewRecord {
		if value == needAdvancedNum {
			return false
		}
	}

	return true
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeRewMax, (*AdvancedRewMaxInfo)(nil))
}
