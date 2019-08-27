package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//升阶消耗返还
type AdvancedExpendReturnInfo struct {
	DanNum    int32                     `json:"danNum"`       //使用进阶丹数量
	RewRecord []int32                   `json:"rewRecord"`    //奖励领取记录
	RewType   welfaretypes.AdvancedType `json:"advancedType"` //系统类型
	IsEmail   bool                      `json:"isEmail"`      //是否邮件
}

func (info *AdvancedExpendReturnInfo) AddRecord(needDanNum int32) {
	info.RewRecord = append(info.RewRecord, needDanNum)
}

func (info *AdvancedExpendReturnInfo) IsCanReceiveRewards(needDanNum int32) bool {
	//条件
	if info.DanNum < needDanNum {
		return false
	}

	//领取记录
	for _, value := range info.RewRecord {
		if value == needDanNum {
			return false
		}
	}

	return true
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeExpendReturn, (*AdvancedExpendReturnInfo)(nil))
}
