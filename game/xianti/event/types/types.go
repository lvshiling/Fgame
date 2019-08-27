package types

import (
	commonlog "fgame/fgame/common/log"
)

type XianTiEventType string

const (
	EventTypeXianTiChanged        XianTiEventType = "XianTiChanged"        //仙体改变事件
	EventTypeXianTiAdvancedCost                   = "XianTiAdvancedCost"   //仙体进阶消耗
	EventTypeXianTiAdvanced                       = "XianTiAdvanced"       //仙体进阶
	EventTypeXianTiUnrealActivate                 = "XianTiUnrealActivate" //激活幻化
	EventTypeXianTiAdvancedLog                    = "XianTiAdvancedLog"    //仙体进阶日志
	EventTypeXianTiPowerChanged                   = "XianTiPowerChanged"   //仙体战力变化
)

//仙体
type PlayerXianTiAdvancedLogEventData struct {
	beforeAdvancedNum int32
	advancedNum       int32
	reason            commonlog.XianTiLogReason
	reasonText        string
}

func CreatePlayerXianTiAdvancedLogEventData(beforeAdvancedNum, advancedNum int32, reason commonlog.XianTiLogReason, reasonText string) *PlayerXianTiAdvancedLogEventData {
	d := &PlayerXianTiAdvancedLogEventData{
		beforeAdvancedNum: beforeAdvancedNum,
		advancedNum:       advancedNum,
		reason:            reason,
		reasonText:        reasonText,
	}
	return d
}

func (d *PlayerXianTiAdvancedLogEventData) GetBeforeAdvancedNum() int32 {
	return d.beforeAdvancedNum
}

func (d *PlayerXianTiAdvancedLogEventData) GetAdvancedNum() int32 {
	return d.advancedNum
}

func (d *PlayerXianTiAdvancedLogEventData) GetReason() commonlog.XianTiLogReason {
	return d.reason
}

func (d *PlayerXianTiAdvancedLogEventData) GetReasonText() string {
	return d.reasonText
}
