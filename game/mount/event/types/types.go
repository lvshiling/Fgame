package types

import (
	commonlog "fgame/fgame/common/log"
)

type MountEventType string

const (
	EventTypeMountChanged        MountEventType = "MountChanged"        //坐骑改变事件
	EventTypeMountAdvancedCost                  = "MountAdvancedCost"   //坐骑进阶消耗
	EventTypeMountAdvanced                      = "MountAdvanced"       //坐骑进阶
	EventTypeMountUnrealActivate                = "MountUnrealActivate" //激活幻化
	EventTypeMountAdvancedLog                   = "MountAdvancedLog"    //坐骑进阶
	EventTypeMountPowerChanged                  = "MountPowerChanged"   //坐骑战力变化
)

//坐骑
type PlayerMountAdvancedLogEventData struct {
	beforeAdvancedNum int32
	advancedNum       int32
	reason            commonlog.MountLogReason
	reasonText        string
}

func CreatePlayerMountAdvancedLogEventData(beforeAdvancedNum, advancedNum int32, reason commonlog.MountLogReason, reasonText string) *PlayerMountAdvancedLogEventData {
	d := &PlayerMountAdvancedLogEventData{
		beforeAdvancedNum: beforeAdvancedNum,
		advancedNum:       advancedNum,
		reason:            reason,
		reasonText:        reasonText,
	}
	return d
}

func (d *PlayerMountAdvancedLogEventData) GetBeforeAdvancedNum() int32 {
	return d.beforeAdvancedNum
}

func (d *PlayerMountAdvancedLogEventData) GetAdvancedNum() int32 {
	return d.advancedNum
}

func (d *PlayerMountAdvancedLogEventData) GetReason() commonlog.MountLogReason {
	return d.reason
}

func (d *PlayerMountAdvancedLogEventData) GetReasonText() string {
	return d.reasonText
}
