package types

import (
	questtypes "fgame/fgame/game/quest/types"
)

type QuestEventType string

const (
	EventTypeQuestFinish QuestEventType = "QuestFinish"
	//任务完成事件
	EventTypeQuestCommit QuestEventType = "QuestCommit"
	//任务开始
	EventTypeQuestAccept QuestEventType = "QuestAccept"
	//一键完成任务类型
	EventTypeQuestFinishAll QuestEventType = "QuestTuMoFinishAll"
)

type QuestLivenessEventType string

const (
	//活跃度过5点
	EventTypeQuestLivenessCrossFive QuestLivenessEventType = "QuestCrossLivenessFive"
)

type QuestDailyEventType string

const (
	//系统帮忙领取日环奖励
	EventTypeQuestDailyReward QuestDailyEventType = "QuestDailyReward"
	//日环任务跨天状态
	EventTypeQuestDailyUpdate QuestDailyEventType = "QuestDailyUpdate"
	//日环过5点
	EventTypeQuestDailyCrossFive QuestDailyEventType = "QuestDailyCrossFive"
)

type QuestKaiFuMuBiaoEventType string

const (
	//开服目标过天
	EventTypeQuestKaiFuMuBiaoCrossDay QuestKaiFuMuBiaoEventType = "QuestKaiFuMuBiaoCrossDay"
	//完成任务次数变更
	EventTypeQuestKaiFuMuBiaoFinishChanged QuestKaiFuMuBiaoEventType = "QuestKaiFuMuBiaoFinishChanged"
)

type QuestFinishAllEventData struct {
	questType questtypes.QuestType
	num       int32
}

func (d *QuestFinishAllEventData) GetQuestType() questtypes.QuestType {
	return d.questType
}

func (d *QuestFinishAllEventData) GetNum() int32 {
	return d.num
}

func CreateQuestFinishAllEventData(questType questtypes.QuestType, num int32) *QuestFinishAllEventData {
	return &QuestFinishAllEventData{
		questType: questType,
		num:       num,
	}
}

// 奇遇任务
type QuestQiYuEventType string

const (
	//奇遇任务结束
	EventTypeQuestQiYuEnd QuestQiYuEventType = "EventTypeQiYuEnd"
)
