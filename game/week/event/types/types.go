package types

import (
	commonlog "fgame/fgame/common/log"
)

type WeekEventType string

const (
	EventTypeWeekBuy WeekEventType = "WeekBuy" //周卡购买
)

//日志事件
type WeekLogEventType string

const (
	EventTypeWeekLogBuy WeekLogEventType = "WeekLogBuy" //周卡购买日志
)

//
type PlayerWeekLogBuyLogEventData struct {
	lastExpireTime int64
	weekType       int32
	reason         commonlog.WeekLogReason
	reasonText     string
}

func CreatePlayerWeekLogBuyLogEventData(lastExpireTime int64, weekType int32, reason commonlog.WeekLogReason, reasonText string) *PlayerWeekLogBuyLogEventData {
	d := &PlayerWeekLogBuyLogEventData{
		lastExpireTime: lastExpireTime,
		weekType:       weekType,
		reason:         reason,
		reasonText:     reasonText,
	}
	return d
}

func (d *PlayerWeekLogBuyLogEventData) GetLastExpireTime() int64 {
	return d.lastExpireTime
}

func (d *PlayerWeekLogBuyLogEventData) GetWeekType() int32 {
	return d.weekType
}

func (d *PlayerWeekLogBuyLogEventData) GetReason() commonlog.WeekLogReason {
	return d.reason
}

func (d *PlayerWeekLogBuyLogEventData) GetReasonText() string {
	return d.reasonText
}
