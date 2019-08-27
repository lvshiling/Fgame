package types

import (
	commonlog "fgame/fgame/common/log"
)

type VipEventType string

const (
	EventTypeVipLevelChanged    VipEventType = "VipLevelChanged"    //vip
	EventTypeVipLevelChangedLog VipEventType = "VipLevelChangedLog" //vip升级日志
)

//
type PlayerVipAdvancedLogEventData struct {
	beforeLevel int32
	beforeGold  int64
	addGold     int64
	reason      commonlog.VipLogReason
	reasonText  string
}

func CreatePlayerVipAdvancedLogEventData(beforeLevel int32, beforeGold, addGold int64, reason commonlog.VipLogReason, reasonText string) *PlayerVipAdvancedLogEventData {
	d := &PlayerVipAdvancedLogEventData{
		beforeLevel: beforeLevel,
		beforeGold:  beforeGold,
		addGold:     addGold,
		reason:      reason,
		reasonText:  reasonText,
	}
	return d
}

func (d *PlayerVipAdvancedLogEventData) GetBeforeLevel() int32 {
	return d.beforeLevel
}

func (d *PlayerVipAdvancedLogEventData) GetBeforeGold() int64 {
	return d.beforeGold
}

func (d *PlayerVipAdvancedLogEventData) GetAddGold() int64 {
	return d.addGold
}

func (d *PlayerVipAdvancedLogEventData) GetReason() commonlog.VipLogReason {
	return d.reason
}

func (d *PlayerVipAdvancedLogEventData) GetReasonText() string {
	return d.reasonText
}
