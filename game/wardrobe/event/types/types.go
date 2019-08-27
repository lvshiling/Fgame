package types

import (
	commonlog "fgame/fgame/common/log"
)

type WardrobeEventType string

const (
	//激活衣橱套装
	EventTypeWardrobeActive WardrobeEventType = "WardrobeActive"
	//带时效性衣橱失效
	EventTypeWardrobeRemove WardrobeEventType = "WardrobeRemove"
	//资质培养升级
	EventTypeWardrobePeiYangUpgrade WardrobeEventType = "WardrobePeiYangUpgrade"
	//日志
	EventTypeWardrobePeiYangLog WardrobeEventType = "WardrobePeiYangUpgradeLog" // 等级升级日志
)

type WardrobePeiYangEventData struct {
	typ      int32
	oldLevel int32
	newLevel int32
}

func (s *WardrobePeiYangEventData) GetType() int32 {
	return s.typ
}

func (s *WardrobePeiYangEventData) GetOldLevel() int32 {
	return s.oldLevel
}

func (s *WardrobePeiYangEventData) GetNewLevel() int32 {
	return s.newLevel
}

func CreateWardrobePeiYangEventData(typ int32, oldLevel int32, newLevel int32) *WardrobePeiYangEventData {
	d := &WardrobePeiYangEventData{
		typ:      typ,
		oldLevel: oldLevel,
		newLevel: newLevel,
	}
	return d
}

//等级升级日志
type WardrobePeiYangLogEventData struct {
	typ        int32
	curLevel   int32
	beforeLev  int32
	reason     commonlog.WardrobeLogReason
	reasonText string
}

func CreateWardrobePeiYangLogEventData(
	typ int32,
	curLevel int32,
	beforeLev int32,
	reason commonlog.WardrobeLogReason,
	reasonText string) *WardrobePeiYangLogEventData {
	d := &WardrobePeiYangLogEventData{
		typ:        typ,
		curLevel:   curLevel,
		beforeLev:  beforeLev,
		reason:     reason,
		reasonText: reasonText,
	}
	return d
}

func (d *WardrobePeiYangLogEventData) GetType() int32 {
	return d.typ
}

func (d *WardrobePeiYangLogEventData) GetCurLevel() int32 {
	return d.curLevel
}

func (d *WardrobePeiYangLogEventData) GetBeforeLev() int32 {
	return d.beforeLev
}

func (d *WardrobePeiYangLogEventData) GetReason() commonlog.WardrobeLogReason {
	return d.reason
}

func (d *WardrobePeiYangLogEventData) GetReasonText() string {
	return d.reasonText
}
