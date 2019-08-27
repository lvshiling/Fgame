package types

import (
	commonlog "fgame/fgame/common/log"
)

type XueDunEventType string

const (
	EventTypeXueDunUpgrade       = "XueDunUpgrade"       //血盾升阶
	EventTypeXueDunNumberChanged = "XueDunNumberChanged" //血盾阶数改变
	EventTypeXueDunBloodChanged  = "XueDunChanged"       //血盾血炼值改变
	EventTypeXueDunUpgradeLog    = "XueDunUpgradeLog "   //血盾升阶日志
)

//血盾
type PlayerXueDunUpgradeLogEventData struct {
	beforeNumber int32
	beforeStar   int32
	number       int32
	star         int32
	reason       commonlog.XueDunLogReason
	reasonText   string
}

func CreatePlayerXueDunUpgradeLogEventData(beforeNumber, beforeStar int32, number int32, star int32, reason commonlog.XueDunLogReason, reasonText string) *PlayerXueDunUpgradeLogEventData {
	d := &PlayerXueDunUpgradeLogEventData{
		beforeNumber: beforeNumber,
		beforeStar:   beforeStar,
		number:       number,
		star:         star,
		reason:       reason,
		reasonText:   reasonText,
	}
	return d
}

func (d *PlayerXueDunUpgradeLogEventData) GetBeforeNumber() int32 {
	return d.beforeNumber
}

func (d *PlayerXueDunUpgradeLogEventData) GetBeforeStar() int32 {
	return d.beforeStar
}

func (d *PlayerXueDunUpgradeLogEventData) GetNumber() int32 {
	return d.number
}

func (d *PlayerXueDunUpgradeLogEventData) GetStar() int32 {
	return d.star
}

func (d *PlayerXueDunUpgradeLogEventData) GetReason() commonlog.XueDunLogReason {
	return d.reason
}

func (d *PlayerXueDunUpgradeLogEventData) GetReasonText() string {
	return d.reasonText
}
