package types

import (
	"fgame/fgame/game/global"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//进阶奖励（随功能开启）
type AdvancedRewExtendedInfo struct {
	RewType     welfaretypes.AdvancedType `json:"rewType"`     //进阶类型
	AdvancedNum int32                     `json:"advancedNum"` //阶数
	ExpireTime  int64                     `json:"expireTime"`  //免费过期时间
	RewRecord   []int32                   `json:"rewRecord"`   //奖励领取记录
	IsEmail     bool                      `json:"isEmail"`
}

func (info *AdvancedRewExtendedInfo) AddRecord(needAdvacedNum int32) {
	info.RewRecord = append(info.RewRecord, needAdvacedNum)
}

func (info *AdvancedRewExtendedInfo) IsCanReceiveRewards(needAdvacedNum int32) bool {
	// 过期时间
	now := global.GetGame().GetTimeService().Now()
	if now > info.ExpireTime {
		return false
	}

	//等级
	if info.AdvancedNum < needAdvacedNum {
		return false
	}

	//领取记录
	for _, value := range info.RewRecord {
		if value == needAdvacedNum {
			return false
		}
	}

	return true
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeRewExtended, (*AdvancedRewExtendedInfo)(nil))
}
