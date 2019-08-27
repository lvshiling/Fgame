package types

import (
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/game/lingtongdev/types"
)

type LingTongDevEventType string

const (
	//灵童养成改变事件
	EventTypeLingTongDevChanged        LingTongDevEventType = "LingTongDevChanged"
	EventTypeLingTongDevUse                                 = "LingTongDevUse"
	EventTypeLingTongDevTrialOverdue                        = "LingTongDevTrialOverdue"
	EventTypeLingTongDevAdvancedCost                        = "LingTongDevAdvancedCost"
	EventTypeLingTongDevAdvanced                            = "LingTongDevAdvanced"       //灵童养成进阶
	EventTypeLingTongDevUnrealActivate                      = "LingTongDevUnrealActivate" //激活幻化
	EventTypeLingTongDevAdvancedLog                         = "LingTongDevAdvancedLog"    //灵童养成进阶日志
	EventTypeLingTongDevPowerChanged                        = "LingTongDevPowerChanged"   //灵童养成战力变化
)

//灵童养成日志
type PlayerLingTongDevAdvancedLogEventData struct {
	classType         types.LingTongDevSysType
	beforeAdvancedNum int32
	advancedNum       int32
	reason            commonlog.LingTongDevLogReason
	reasonText        string
}

func CreatePlayerLingTongDevAdvancedLogEventData(classType types.LingTongDevSysType, beforeAdvancedNum, advancedNum int32, reason commonlog.LingTongDevLogReason, reasonText string) *PlayerLingTongDevAdvancedLogEventData {
	d := &PlayerLingTongDevAdvancedLogEventData{
		classType:         classType,
		beforeAdvancedNum: beforeAdvancedNum,
		advancedNum:       advancedNum,
		reason:            reason,
		reasonText:        reasonText,
	}
	return d
}

func (d *PlayerLingTongDevAdvancedLogEventData) GetClassType() types.LingTongDevSysType {
	return d.classType
}

func (d *PlayerLingTongDevAdvancedLogEventData) GetBeforeAdvancedNum() int32 {
	return d.beforeAdvancedNum
}

func (d *PlayerLingTongDevAdvancedLogEventData) GetAdvancedNum() int32 {
	return d.advancedNum
}

func (d *PlayerLingTongDevAdvancedLogEventData) GetReason() commonlog.LingTongDevLogReason {
	return d.reason
}

func (d *PlayerLingTongDevAdvancedLogEventData) GetReasonText() string {
	return d.reasonText
}
