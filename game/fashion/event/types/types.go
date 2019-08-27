package types

import (
	fashiontypes "fgame/fgame/game/fashion/types"
)

type FashionEventType string

const (
	//时装改变事件
	EventTypeFashionChanged FashionEventType = "FashionChanged"
	//时装过期
	EventTypeFashionOverdue FashionEventType = "FashionOverdue"
	//时装激活
	EventTypeFashionActivate FashionEventType = "FashionActivate"
	//时装试用过期
	EventTypeFashionTrialOverdue FashionEventType = "FashionTrialOverdue"
)

type FashionTrialOverdueEventData struct {
	trialId     int32
	overdueType fashiontypes.FashionTrialOverdueType
}

func CreateFashionTrialOverdueEventData(trialId int32, overdueType fashiontypes.FashionTrialOverdueType) *FashionTrialOverdueEventData {
	d := &FashionTrialOverdueEventData{
		trialId:     trialId,
		overdueType: overdueType,
	}
	return d
}

func (w *FashionTrialOverdueEventData) GetTrialId() int32 {
	return w.trialId
}

func (w *FashionTrialOverdueEventData) GetOverdueType() fashiontypes.FashionTrialOverdueType {
	return w.overdueType
}
