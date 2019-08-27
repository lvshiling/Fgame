package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//返利-充值培养
type FeedbackDevelopInfo struct {
	ActivateChargeNum int64 `json:"activateChargeNum"` //激活充值
	IsDead            bool  `json:"isDead"`            //是否死亡
	IsFeed            bool  `json:"isFeed"`            //是否喂养
	IsActivate        bool  `json:"isActivate"`        //是否激活
	IsReceiveRew      bool  `json:"isReceiveRew"`      //是否领取最终奖励
	TodayCostNum      int64 `json:"todayCostNum"`      //今日消费
	FeedTimes         int32 `json:"feedDay"`           //累计喂养天数
	IsEndMail         bool  `json:"isEndMail"`         //结束邮件
}

//是否可以复活金鸡
func (info *FeedbackDevelopInfo) IsCanRevive() bool {
	if info.IsActivate && !info.IsDead {
		return false
	}

	return true
}

//当天喂养奖励
func (info *FeedbackDevelopInfo) IsCanReceiveToday(needGost int64) bool {
	if info.IsFeed {
		return false
	}

	// 喂养条件
	if info.TodayCostNum < needGost {
		return false
	}

	return true
}

//累计充值奖励
func (info *FeedbackDevelopInfo) IsCanReceiveCountDay(needTimes int32) bool {
	//记录
	if info.IsReceiveRew {
		return false
	}

	//条件
	if info.FeedTimes < needTimes {
		return false
	}

	return true
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeDevelop, (*FeedbackDevelopInfo)(nil))
}
