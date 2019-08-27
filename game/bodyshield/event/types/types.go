package types

import (
	commonlog "fgame/fgame/common/log"
)

type BodyShieldEventType string

const (
	EventTypeBodyShieldAdvancedCost BodyShieldEventType = "BodyShieldAdvancedCost" //护体盾进阶消耗
	EventTypeShieldAdvancedCost                         = "ShieldAdvancedCost"     //盾刺进阶消耗
	EventTypeBodyShieldAdvanced                         = "BodyShieldAdvanced"     //护体盾进阶
	EventTypeShieldAdvanced                             = "ShieldAdvanced"         //神盾尖刺进阶
	EventTypeBodyShieldAdvancedLog                      = "BodyShieldAdvancedLog"  //护体盾进阶日志
	EventTypeShieldAdvancedLog                          = "ShieldAdvancedLog"      //神盾尖刺进阶日志
	EventTypeShieldPowerChanged                         = "ShieldPowerChanged"     //神盾尖刺战力变化
	EventTypeBodyShieldPowerChanged                     = "BodyShieldPowerChanged" //护体盾战力变化
)

//
type PlayerBodyShieldAdvancedLogEventData struct {
	beforeAdvancedNum int32
	advancedNum       int32
	reason            commonlog.ShieldLogReason
	reasonText        string
}

func CreatePlayerBodyShieldAdvancedLogEventData(beforeAdvancedNum, advancedNum int32, reason commonlog.ShieldLogReason, reasonText string) *PlayerBodyShieldAdvancedLogEventData {
	d := &PlayerBodyShieldAdvancedLogEventData{
		beforeAdvancedNum: beforeAdvancedNum,
		advancedNum:       advancedNum,
		reason:            reason,
		reasonText:        reasonText,
	}
	return d
}

func (d *PlayerBodyShieldAdvancedLogEventData) GetBeforeAdvancedNum() int32 {
	return d.beforeAdvancedNum
}

func (d *PlayerBodyShieldAdvancedLogEventData) GetAdvancedNum() int32 {
	return d.advancedNum
}

func (d *PlayerBodyShieldAdvancedLogEventData) GetReason() commonlog.ShieldLogReason {
	return d.reason
}

func (d *PlayerBodyShieldAdvancedLogEventData) GetReasonText() string {
	return d.reasonText
}

//盾刺
type PlayerShieldAdvancedLogEventData struct {
	beforeAdvancedNum int32
	advancedNum       int32
	reason            commonlog.ShieldLogReason
	reasonText        string
}

func CreatePlayerShieldAdvancedLogEventData(beforeAdvancedNum, advancedNum int32, reason commonlog.ShieldLogReason, reasonText string) *PlayerShieldAdvancedLogEventData {
	d := &PlayerShieldAdvancedLogEventData{
		beforeAdvancedNum: beforeAdvancedNum,
		advancedNum:       advancedNum,
		reason:            reason,
		reasonText:        reasonText,
	}
	return d
}

func (d *PlayerShieldAdvancedLogEventData) GetBeforeAdvancedNum() int32 {
	return d.beforeAdvancedNum
}

func (d *PlayerShieldAdvancedLogEventData) GetAdvancedNum() int32 {
	return d.advancedNum
}

func (d *PlayerShieldAdvancedLogEventData) GetReason() commonlog.ShieldLogReason {
	return d.reason
}

func (d *PlayerShieldAdvancedLogEventData) GetReasonText() string {
	return d.reasonText
}
