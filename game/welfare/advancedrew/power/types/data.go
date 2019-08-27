package types

import (
	"fgame/fgame/game/global"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//进阶战力奖励（随功能开启）
type AdvancedPowerInfo struct {
	RewType    welfaretypes.AdvancedType `json:"rewType"`    //进阶类型
	Power      int64                     `json:"power"`      //战力
	ExpireTime int64                     `json:"expireTime"` //免费过期时间
	RewRecord  []int64                   `json:"rewRecord"`  //奖励领取记录
	IsEmail    bool                      `json:"isEmail"`
}

func (info *AdvancedPowerInfo) AddRecord(needPowerNum int64) {
	info.RewRecord = append(info.RewRecord, needPowerNum)
}

func (info *AdvancedPowerInfo) IsCanReceiveRewards(needPowerNum int64) bool {
	// 过期时间
	now := global.GetGame().GetTimeService().Now()
	if now > info.ExpireTime {
		return false
	}

	// 等级
	if info.Power < needPowerNum {
		return false
	}
	//领取记录
	for _, value := range info.RewRecord {
		if value == needPowerNum {
			return false
		}
	}

	return true
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypePower, (*AdvancedPowerInfo)(nil))
}
