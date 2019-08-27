package types

import (
	commonlog "fgame/fgame/common/log"
	xianfutypes "fgame/fgame/game/xianfu/types"
)

type XianFuEventType string

const (
	//仙府挑战事件
	EventTypeXianFuChallenge XianFuEventType = "XianFuChallenge"
	//仙府挑战成功
	EventTypeXianFuFinish XianFuEventType = "XianFuFinish"
	//开始升级仙府
	EventTypeXianFuStartUpgrade XianFuEventType = "XianFuStartUpgrade"
	//仙府扫荡事件
	EventTypeXianFuSweep XianFuEventType = "XianFuSweep"
	//仙府升级成功
	EventTypeXianFuUpgradeSuccess XianFuEventType = "XianFuUpgradeSuccess"
	//仙府刷新怪
	EventTypeXianFuRefreshGroup XianFuEventType = "XianFuRefreshGroup"
	//后台仙府日志
	EventTypeXianFuLog XianFuEventType = "XianFuLog"
)

type XianFuChallengeEventData struct {
	typ xianfutypes.XianfuType
	num int32
}

func (d *XianFuChallengeEventData) GetType() xianfutypes.XianfuType {
	return d.typ
}

func (d *XianFuChallengeEventData) GetNum() int32 {
	return d.num
}

func CreateXianFuChallengeEventData(typ xianfutypes.XianfuType, num int32) *XianFuChallengeEventData {
	return &XianFuChallengeEventData{
		typ: typ,
		num: num,
	}
}

type XianFuFinishEventData struct {
	typ xianfutypes.XianfuType
	num int32
}

func (d *XianFuFinishEventData) GetType() xianfutypes.XianfuType {
	return d.typ
}

func (d *XianFuFinishEventData) GetNum() int32 {
	return d.num
}

func CreateXianFuFinishEventData(typ xianfutypes.XianfuType, num int32) *XianFuFinishEventData {
	return &XianFuFinishEventData{
		typ: typ,
		num: num,
	}
}

// 日志
type PlayerXianFuLogEventData struct {
	beforeLevel int32
	uplevel     int32
	xianfuType  xianfutypes.XianfuType
	reason      commonlog.XianFuLogReason
	reasonText  string
}

func CreatePlayerXianFuLogEventData(beforeLevel, uplevel int32, typ xianfutypes.XianfuType, reason commonlog.XianFuLogReason, reasonText string) *PlayerXianFuLogEventData {
	d := &PlayerXianFuLogEventData{
		beforeLevel: beforeLevel,
		uplevel:     uplevel,
		xianfuType:  typ,
		reason:      reason,
		reasonText:  reasonText,
	}
	return d
}

func (d *PlayerXianFuLogEventData) GetBeforeLevel() int32 {
	return d.beforeLevel
}

func (d *PlayerXianFuLogEventData) GetUplevel() int32 {
	return d.uplevel
}

func (d *PlayerXianFuLogEventData) GetXianFuType() xianfutypes.XianfuType {
	return d.xianfuType
}

func (d *PlayerXianFuLogEventData) GetReason() commonlog.XianFuLogReason {
	return d.reason
}

func (d *PlayerXianFuLogEventData) GetReasonText() string {
	return d.reasonText
}
