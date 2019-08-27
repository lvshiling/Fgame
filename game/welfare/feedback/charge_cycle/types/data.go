package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//返利-连续充值
type FeedbackCycleChargeInfo struct {
	CurDayChargeNum int32   `json:"curDayChargeNum"` //今日充值数量
	IsReceiveDayRew bool    `json:"isReceiveDayRew"` //当天奖励是否领取
	DayNum          int32   `json:"dayNum"`          //累计充值天数
	RewRecord       []int32 `json:"rewRecord"`       //累计奖励领取记录
	CycleDay        int32   `json:"cycleDay"`        //第几天的充值数据
	IsEndMail       bool    `json:"isEndMail"`       //结束邮件
}

// //当天充值目标
// func (info *FeedbackCycleChargeInfo) AddDayCharge(addGold, needGold int32) {
// 	cur := info.CurDayChargeNum
// 	lessBefor := cur < needGold
// 	cur += addGold
// 	greaterAfter := cur >= needGold
// 	if lessBefor && greaterAfter {
// 		info.DayNum += 1
// 	}

// 	info.CurDayChargeNum = cur
// }

//当天充值目标
func (info *FeedbackCycleChargeInfo) UpdateTodayCharge(todayChargeNum, needGold int32) {
	cur := info.CurDayChargeNum
	lessBefor := cur < needGold
	greaterAfter := todayChargeNum >= needGold
	if lessBefor && greaterAfter {
		info.DayNum += 1
	}
	info.CurDayChargeNum = todayChargeNum
}

//当天充值奖励
func (info *FeedbackCycleChargeInfo) IsCanReceiveToday(needGold int32) bool {
	if info.IsReceiveDayRew {
		return false
	}

	// 充值条件
	if info.CurDayChargeNum < needGold {
		return false
	}

	return true
}

//累计充值奖励
func (info *FeedbackCycleChargeInfo) IsCanReceiveCountDay(needDay int32) bool {
	//条件
	if info.DayNum < needDay {
		return false
	}

	//领取记录
	for _, value := range info.RewRecord {
		if value == needDay {
			return false
		}
	}
	return true
}

//领取当天充值奖励
func (info *FeedbackCycleChargeInfo) ReceiveToday() {
	info.IsReceiveDayRew = true
}

//领取累计充值奖励
func (info *FeedbackCycleChargeInfo) ReceiveCountDay(needDay int32) {
	info.RewRecord = append(info.RewRecord, needDay)
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeCycleCharge, (*FeedbackCycleChargeInfo)(nil))
}
