package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 活动目标
type GoalInfo struct {
	GoalCount    int32              `json:"goalCount"`    //参与次数
	RewRecordMap map[int32]struct{} `json:"rewRecordMap"` //奖励领取记录
	IsEmail      bool               `json:"isEmail"`      //是否邮件
}

func (info *GoalInfo) ReachGoal() {
	info.GoalCount += 1
}

func (info *GoalInfo) AddRecord(goalCount int32) {
	info.RewRecordMap[goalCount] = struct{}{}
}

func (info *GoalInfo) GetRewRecord() (recordList []int32) {
	for rewCount, _ := range info.RewRecordMap {
		recordList = append(recordList, rewCount)
	}

	return
}

func (info *GoalInfo) IsCanReceiveRewards(rewGoalCount int32) bool {
	if info.GoalCount < rewGoalCount {
		return false
	}

	//领取记录
	_, ok := info.RewRecordMap[rewGoalCount]
	if ok {
		return false
	}

	return true
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeHuHu, welfaretypes.OpenActivitySpecialSubTypeGoal, (*GoalInfo)(nil))
}
