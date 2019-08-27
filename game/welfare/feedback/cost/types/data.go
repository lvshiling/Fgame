package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//返利-消费
type FeedbackCostInfo struct {
	GoldNum   int32   `json:"goldNum"`   //消费元宝数量
	RewRecord []int32 `json:"rewRecord"` //奖励领取记录
	IsEmail   bool    `json:"isEmail"`   //是否发邮件
}

func (info *FeedbackCostInfo) IsCanReceiveRewards(goldNum int32) bool {
	//条件
	if info.GoldNum < goldNum {
		return false
	}

	//领取记录
	for _, value := range info.RewRecord {
		if value == goldNum {
			return false
		}
	}

	return true
}

func (info *FeedbackCostInfo) AddRecord(goldNum int32) {
	info.RewRecord = append(info.RewRecord, goldNum)
}
func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeCost, (*FeedbackCostInfo)(nil))
}
