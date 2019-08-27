package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	
)

//祝福丹大放送
type BlessAdvancedInfo struct {
	AdvancedNum int32   `json:"advancedNum"` //阶数
	RewRecord   []int32 `json:"rewRecord"`   //奖励领取记录
	BlessDay    int32   `json:"blessDay"`    //当前奖励日
}

func (info *BlessAdvancedInfo) AddRecord(needAdvancedNum int32) {
	info.RewRecord = append(info.RewRecord, needAdvancedNum)
}

func (info *BlessAdvancedInfo) IsCanReceiveRewards(needAdvancedNum int32) bool {
	if info.AdvancedNum < needAdvancedNum {
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
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeAdvanced, welfaretypes.OpenActivityAdvancedSubTypeBlessFeedback, (*BlessAdvancedInfo)(nil))
}
