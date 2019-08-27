package types

import (
	commonlog "fgame/fgame/common/log"
)

type ShenfaEventType string

const (
	EventTypeShenfaChanged        ShenfaEventType = "ShenfaChanged" //身法改变事件
	EventTypeShenfaUse                            = "ShenfaUse"
	EventTypeShenfaAdvanced                       = "ShenfaAdvanced"       //身法进阶事件
	EventTypeShenfaAdvancedCost                   = "ShenfaAdvancedCost"   //身法进阶消耗
	EventTypeShenfaAdvancedLog                    = "ShenfaAdvancedLog"    //身法进阶事件
	EventTypeShenfaPowerChanged                   = "ShenfaPowerChanged"   //身法战力变化
	EventTypeShenfaUnrealActivate                 = "ShenfaUnrealActivate" //激活幻化
)

//身法
type PlayerShenfaAdvancedLogEventData struct {
	beforeAdvancedNum int32
	advancedNum       int32
	reason            commonlog.ShenfaLogReason
	reasonText        string
}

func CreatePlayerShenfaAdvancedLogEventData(beforeAdvancedNum, advancedNum int32, reason commonlog.ShenfaLogReason, reasonText string) *PlayerShenfaAdvancedLogEventData {
	d := &PlayerShenfaAdvancedLogEventData{
		beforeAdvancedNum: beforeAdvancedNum,
		advancedNum:       advancedNum,
		reason:            reason,
		reasonText:        reasonText,
	}
	return d
}

func (d *PlayerShenfaAdvancedLogEventData) GetBeforeAdvancedNum() int32 {
	return d.beforeAdvancedNum
}

func (d *PlayerShenfaAdvancedLogEventData) GetAdvancedNum() int32 {
	return d.advancedNum
}

func (d *PlayerShenfaAdvancedLogEventData) GetReason() commonlog.ShenfaLogReason {
	return d.reason
}

func (d *PlayerShenfaAdvancedLogEventData) GetReasonText() string {
	return d.reasonText
}
