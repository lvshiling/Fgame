package types

import (
	commonlog "fgame/fgame/common/log"
)

type WingEventType string

const (
	//战翼改变事件
	EventTypeWingChanged         WingEventType = "WingChanged"
	EventTypeWingUse                           = "WingUse"
	EventTypeWingTrialOverdue                  = "WingTrialOverdue"
	EventTypeWingAdvancedCost                  = "WingAdvancedCost"
	EventTypeWingAdvanced                      = "WingAdvanced"        //战翼进阶
	EventTypeWingUnrealActivate                = "WingUnrealActivate"  //激活幻化
	EventTypeFeatherAdvanced                   = "FeatherAdvanced"     //护体仙羽进阶
	EventTypeWingAdvancedLog                   = "WingAdvancedLog"     //战翼进阶日志
	EventTypeFeatherAdvancedLog                = "FeatherAdvancedLog"  //护体仙羽进阶日志
	EventTypeFeatherPowerChanged               = "FeatherPowerChanged" //护体仙羽战力变化
	EventTypeWingPowerChanged                  = "WingPowerChanged"    //战翼战力变化
)

type WingTrialOverdueEventData struct {
	trialId int32
	bResult bool
}

func CreateWingTrialOverdueEventData(trialId int32, bResult bool) *WingTrialOverdueEventData {
	d := &WingTrialOverdueEventData{
		trialId: trialId,
		bResult: bResult,
	}
	return d
}

func (w *WingTrialOverdueEventData) GetTrialId() int32 {
	return w.trialId
}

func (w *WingTrialOverdueEventData) GetResult() bool {
	return w.bResult
}

//战翼日志
type PlayerWingAdvancedLogEventData struct {
	beforeAdvancedNum int32
	advancedNum       int32
	reason            commonlog.WingLogReason
	reasonText        string
}

func CreatePlayerWingAdvancedLogEventData(beforeAdvancedNum, advancedNum int32, reason commonlog.WingLogReason, reasonText string) *PlayerWingAdvancedLogEventData {
	d := &PlayerWingAdvancedLogEventData{
		beforeAdvancedNum: beforeAdvancedNum,
		advancedNum:       advancedNum,
		reason:            reason,
		reasonText:        reasonText,
	}
	return d
}

func (d *PlayerWingAdvancedLogEventData) GetBeforeAdvancedNum() int32 {
	return d.beforeAdvancedNum
}

func (d *PlayerWingAdvancedLogEventData) GetAdvancedNum() int32 {
	return d.advancedNum
}

func (d *PlayerWingAdvancedLogEventData) GetReason() commonlog.WingLogReason {
	return d.reason
}

func (d *PlayerWingAdvancedLogEventData) GetReasonText() string {
	return d.reasonText
}

//仙羽日志
type PlayerFeatherAdvancedLogEventData struct {
	beforeAdvancedNum int32
	advancedNum       int32
	reason            commonlog.WingLogReason
	reasonText        string
}

func CreatePlayerFeatherAdvancedLogEventData(beforeAdvancedNum, advancedNum int32, reason commonlog.WingLogReason, reasonText string) *PlayerFeatherAdvancedLogEventData {
	d := &PlayerFeatherAdvancedLogEventData{
		beforeAdvancedNum: beforeAdvancedNum,
		advancedNum:       advancedNum,
		reason:            reason,
		reasonText:        reasonText,
	}
	return d
}

func (d *PlayerFeatherAdvancedLogEventData) GetBeforeAdvancedNum() int32 {
	return d.beforeAdvancedNum
}

func (d *PlayerFeatherAdvancedLogEventData) GetAdvancedNum() int32 {
	return d.advancedNum
}

func (d *PlayerFeatherAdvancedLogEventData) GetReason() commonlog.WingLogReason {
	return d.reason
}

func (d *PlayerFeatherAdvancedLogEventData) GetReasonText() string {
	return d.reasonText
}
