package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//返利-单笔充值
type FeedbackSingleChargeInfo struct {
	MaxSingleChargeNum int32   `json:"maxSingleChargeNum"` //单笔最大数
	RewRecord          []int32 `json:"rewRecord"`          //领取记录
	IsEmail            bool    `json:"isEmail"`            //是否奖励发放
}

func (info *FeedbackSingleChargeInfo) IsCanReceiveRewards(needGoldNum int32) bool {
	//条件
	if info.MaxSingleChargeNum < needGoldNum {
		return false
	}

	//领取记录
	for _, value := range info.RewRecord {
		if value == needGoldNum {
			return false
		}
	}

	return true
}

func (info *FeedbackSingleChargeInfo) AddRecord(needGoldNum int32) {
	info.RewRecord = append(info.RewRecord, needGoldNum)
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeSingleChagre, (*FeedbackSingleChargeInfo)(nil))
}
