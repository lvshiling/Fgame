package types

import (
	commonlog "fgame/fgame/common/log"
	tianshutypes "fgame/fgame/game/tianshu/types"
)

type TianShuEventType string

var (
	EventTypePlayerTianShuActivate TianShuEventType = "PlayerTianShuActivate" //天书激活
	EventTypePlayerTianShuUplevel                   = "PlayerTianShuUplevel"  //天书升级
	EventTypePlayerTianShuLog                       = "PlayerTianShuLog"      //天书日志
)

//
type PlayerTianShuLogEventData struct {
	typ         tianshutypes.TianShuType
	beforeLevel int32
	uplevel     int32
	reason      commonlog.TianShuLogReason
	reasonText  string
}

func CreatePlayerTianShuLogEventData(beforLevel, uplevel int32, typ tianshutypes.TianShuType, reason commonlog.TianShuLogReason, reasonText string) *PlayerTianShuLogEventData {
	d := &PlayerTianShuLogEventData{
		typ:         typ,
		beforeLevel: beforLevel,
		uplevel:     uplevel,
		reason:      reason,
		reasonText:  reasonText,
	}
	return d
}

func (d *PlayerTianShuLogEventData) GetTianShuType() tianshutypes.TianShuType {
	return d.typ
}

func (d *PlayerTianShuLogEventData) GetBeforeLevel() int32 {
	return d.beforeLevel
}

func (d *PlayerTianShuLogEventData) GetUpLevel() int32 {
	return d.uplevel
}

func (d *PlayerTianShuLogEventData) GetReason() commonlog.TianShuLogReason {
	return d.reason
}

func (d *PlayerTianShuLogEventData) GetReasonText() string {
	return d.reasonText
}
