package types

import (
	lingtongtypes "fgame/fgame/game/lingtong/types"
)

type LingTongFashionEventType string

const (
	//时装改变事件
	EventTypeLingTongFashionChanged LingTongFashionEventType = "LingTongFashionChanged"
	//时装过期
	EventTypeLingTongFashionOverdue LingTongFashionEventType = "LignTongFashionOverdue"
	//时装激活
	EventTypeLingTongFashionActivate LingTongFashionEventType = "LingTongFashionActivate"
	//时装试用过期
	EventTypeLingTongFashionTrialOverdue LingTongFashionEventType = "LingTongFashionTrialOverdue"
)

type LingTongFashionTrialOverdueEventData struct {
	trialId     int32
	overdueType lingtongtypes.LingTongFashionTrialOverdueType
}

func CreateLingTongFashionTrialOverdueEventData(trialId int32, overdueType lingtongtypes.LingTongFashionTrialOverdueType) *LingTongFashionTrialOverdueEventData {
	d := &LingTongFashionTrialOverdueEventData{
		trialId:     trialId,
		overdueType: overdueType,
	}
	return d
}

func (w *LingTongFashionTrialOverdueEventData) GetTrialId() int32 {
	return w.trialId
}

func (w *LingTongFashionTrialOverdueEventData) GetOverdueType() lingtongtypes.LingTongFashionTrialOverdueType {
	return w.overdueType
}
