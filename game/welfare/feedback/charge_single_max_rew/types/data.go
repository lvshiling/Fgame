package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//返利-单笔充值(最近档次)
type FeedbackSingleChargeMaxRewInfo struct {
	MaxSingleChargeNum int32   `json:"maxSingleChargeNum"` //单笔最大数
	CanRewRecord       []int32 `json:"canRewRecord"`       //可领取记录
	ReceiveRewRecord   []int32 `json:"receiveRewRecord"`   //已领取记录
	IsEmail            bool    `json:"isEmail"`            //是否奖励发放
}

func (info *FeedbackSingleChargeMaxRewInfo) IsCanReceiveRewards(needGoldNum int32) bool {
	//条件
	flag := false
	for _, value := range info.CanRewRecord {
		if value != needGoldNum {
			continue
		}
		flag = true
	}

	return flag
}

func (info *FeedbackSingleChargeMaxRewInfo) AddReceiveRecord(needGoldNum int32) {

	//删除 可领取记录
	delIndex := -1
	for index, value := range info.CanRewRecord {
		if value == needGoldNum {
			delIndex = index
			break
		}
	}
	if delIndex == -1 {
		return
	}
	info.CanRewRecord = append(info.CanRewRecord[:delIndex], info.CanRewRecord[delIndex+1:]...)

	// 添加 已领取记录
	info.ReceiveRewRecord = append(info.ReceiveRewRecord, needGoldNum)
}

func (info *FeedbackSingleChargeMaxRewInfo) AddCanRewRecord(needGoldNum int32) {
	info.CanRewRecord = append(info.CanRewRecord, needGoldNum)
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeSingleChagreMaxRew, (*FeedbackSingleChargeMaxRewInfo)(nil))
}
