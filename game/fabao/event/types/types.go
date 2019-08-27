package types

import (
	commonlog "fgame/fgame/common/log"
)

type FaBaoEventType string

const (
	//法宝改变事件
	EventTypeFaBaoChanged        FaBaoEventType = "FaBaoChanged"
	EventTypeFaBaoUse                           = "FaBaoUse"
	EventTypeFaBaoTrialOverdue                  = "FaBaoTrialOverdue"
	EventTypeFaBaoAdvancedCost                  = "FaBaoAdvancedCost"
	EventTypeFaBaoAdvanced                      = "FaBaoAdvanced"       //法宝进阶
	EventTypeFaBaoUnrealActivate                = "FaBaoUnrealActivate" //激活幻化
	EventTypeFaBaoAdvancedLog                   = "FaBaoAdvancedLog"    //法宝进阶日志
	EventTypeFaBaoPowerChanged                  = "FaBaoPowerChanged"   //法宝战力变化
)

//法宝日志
type PlayerFaBaoAdvancedLogEventData struct {
	beforeAdvancedNum int32
	advancedNum       int32
	reason            commonlog.FaBaoLogReason
	reasonText        string
}

func CreatePlayerFaBaoAdvancedLogEventData(beforeAdvancedNum, advancedNum int32, reason commonlog.FaBaoLogReason, reasonText string) *PlayerFaBaoAdvancedLogEventData {
	d := &PlayerFaBaoAdvancedLogEventData{
		beforeAdvancedNum: beforeAdvancedNum,
		advancedNum:       advancedNum,
		reason:            reason,
		reasonText:        reasonText,
	}
	return d
}

func (d *PlayerFaBaoAdvancedLogEventData) GetBeforeAdvancedNum() int32 {
	return d.beforeAdvancedNum
}

func (d *PlayerFaBaoAdvancedLogEventData) GetAdvancedNum() int32 {
	return d.advancedNum
}

func (d *PlayerFaBaoAdvancedLogEventData) GetReason() commonlog.FaBaoLogReason {
	return d.reason
}

func (d *PlayerFaBaoAdvancedLogEventData) GetReasonText() string {
	return d.reasonText
}
