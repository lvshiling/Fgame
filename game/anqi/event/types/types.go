package types

import (
	commonlog "fgame/fgame/common/log"
)

type AnqiEventType string

const (
	EventTypeAnqiAdvanced     AnqiEventType = "AnqiAdvanced"     //暗器进阶事件
	EventTypeAnqiAdvancedCost               = "AnqiAdvancedCost" //暗器进阶消耗
	EventTypeAnqiAdvancedLog                = "AnqiAdvancedLog"  //暗器进阶日志
	EventTypeAnqiPowerChanged               = "AnqiPowerChanged" //暗器战力变化
)

//
type PlayerAnqiAdvancedLogEventData struct {
	beforeAdvancedNum int32
	advancedNum       int32
	reason            commonlog.AnqiLogReason
	reasonText        string
}

func CreatePlayerAnqiAdvancedLogEventData(beforeAdvancedNum, advancedNum int32, reason commonlog.AnqiLogReason, reasonText string) *PlayerAnqiAdvancedLogEventData {
	d := &PlayerAnqiAdvancedLogEventData{
		beforeAdvancedNum: beforeAdvancedNum,
		advancedNum:       advancedNum,
		reason:            reason,
		reasonText:        reasonText,
	}
	return d
}

func (d *PlayerAnqiAdvancedLogEventData) GetBeforeAdvancedNum() int32 {
	return d.beforeAdvancedNum
}

func (d *PlayerAnqiAdvancedLogEventData) GetAdvancedNum() int32 {
	return d.advancedNum
}

func (d *PlayerAnqiAdvancedLogEventData) GetReason() commonlog.AnqiLogReason {
	return d.reason
}

func (d *PlayerAnqiAdvancedLogEventData) GetReasonText() string {
	return d.reasonText
}
