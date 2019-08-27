package types

import (
	commonlog "fgame/fgame/common/log"
)

type LingyuEventType string

const (
	EventTypeLingyuChanged        LingyuEventType = "LingyuChanged"  //领域改变事件(场景改变ui)
	EventTypeLingyuAdvanced                       = "LingyuAdvanced" //领域进阶事件
	EventTypeLingyuUse                            = "LingyuUse"
	EventTypeLingyuSkillChanged                   = "LingyuSkillChanged"   //领域技能替换
	EventTypeLingyuAdvancedCost                   = "LingyuAdvancedCost"   //身法进阶消耗
	EventTypeLingyuAdvancedLog                    = "LingyuAdvancedLog"    //领域进阶日志
	EventTypeLingyuPowerChanged                   = "LingyuPowerChanged"   //领域战力变化
	EventTypeLingyuUnrealActivate                 = "LingyuUnrealActivate" //激活幻化
)

//领域
type PlayerLingyuAdvancedLogEventData struct {
	beforeAdvancedNum int32
	advancedNum       int32
	reason            commonlog.LingyuLogReason
	reasonText        string
}

func CreatePlayerLingyuAdvancedLogEventData(beforeAdvancedNum, advancedNum int32, reason commonlog.LingyuLogReason, reasonText string) *PlayerLingyuAdvancedLogEventData {
	d := &PlayerLingyuAdvancedLogEventData{
		beforeAdvancedNum: beforeAdvancedNum,
		advancedNum:       advancedNum,
		reason:            reason,
		reasonText:        reasonText,
	}
	return d
}

func (d *PlayerLingyuAdvancedLogEventData) GetBeforeAdvancedNum() int32 {
	return d.beforeAdvancedNum
}

func (d *PlayerLingyuAdvancedLogEventData) GetAdvancedNum() int32 {
	return d.advancedNum
}

func (d *PlayerLingyuAdvancedLogEventData) GetReason() commonlog.LingyuLogReason {
	return d.reason
}

func (d *PlayerLingyuAdvancedLogEventData) GetReasonText() string {
	return d.reasonText
}
