package types

import (
	commonlog "fgame/fgame/common/log"
)

type FeiShengEventType string

const (
	EventTypePlayerSanGong     FeiShengEventType = "PlayerSanGong"     //玩家散功
	EventTypePlayerFeiSheng                      = "PlayerFeiSheng"    //玩家飞升
	EventTypePlayerFeiShengLog                   = "PlayerFeiShengLog" //玩家飞升日志
)

//飞升日志
type PlayerFeiShengLogEventData struct {
	beforeFeiShengLevel int32
	beforeGongDeNum     int64
	reason              commonlog.FeiShengLogReason
	reasonText          string
}

func CreatePlayerFeiShengLogEventData(beforeFeiShengLevel int32, beforeGongDeNum int64, reason commonlog.FeiShengLogReason, reasonText string) *PlayerFeiShengLogEventData {
	d := &PlayerFeiShengLogEventData{
		beforeFeiShengLevel: beforeFeiShengLevel,
		beforeGongDeNum:     beforeGongDeNum,
		reason:              reason,
		reasonText:          reasonText,
	}
	return d
}

func (d *PlayerFeiShengLogEventData) GetBeforeLevel() int32 {
	return d.beforeFeiShengLevel
}

func (d *PlayerFeiShengLogEventData) GetBeforeGongDe() int64 {
	return d.beforeGongDeNum
}

func (d *PlayerFeiShengLogEventData) GetReason() commonlog.FeiShengLogReason {
	return d.reason
}

func (d *PlayerFeiShengLogEventData) GetReasonText() string {
	return d.reasonText
}
