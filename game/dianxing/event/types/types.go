package types

import (
	commonlog "fgame/fgame/common/log"
)

type DianXingEventType string

const (
	EventTypeDianXingAdvanced           DianXingEventType = "DianXingAdvanced"           //点星系统进阶事件
	EventTypeDianXingAdvancedLog                          = "DianXingAdvancedLog"        //点星系统进阶日志
	EventTypeDianXingJieFengAdvancedLog                   = "DianXingJieFengAdvancedLog" //点星系统解封升级日志
	EventTypeDianXingUseItem                              = "DianXingUseItem"            //点星系统使用星尘
)

//点星系统升级日志
type PlayerDianXingAdvancedLogEventData struct {
	beforeXingPu int32
	beforeLev    int32
	reason       commonlog.DianXingLogReason
	reasonText   string
}

func CreatePlayerDianXingAdvancedLogEventData(beforeXingPu int32, beforeLev int32, reason commonlog.DianXingLogReason, reasonText string) *PlayerDianXingAdvancedLogEventData {
	d := &PlayerDianXingAdvancedLogEventData{
		beforeXingPu: beforeXingPu,
		beforeLev:    beforeLev,
		reason:       reason,
		reasonText:   reasonText,
	}
	return d
}

func (d *PlayerDianXingAdvancedLogEventData) GetBeforeXingPu() int32 {
	return d.beforeXingPu
}

func (d *PlayerDianXingAdvancedLogEventData) GetBeforeLev() int32 {
	return d.beforeLev
}

func (d *PlayerDianXingAdvancedLogEventData) GetReason() commonlog.DianXingLogReason {
	return d.reason
}

func (d *PlayerDianXingAdvancedLogEventData) GetReasonText() string {
	return d.reasonText
}

//点星解封升级日志
type PlayerDianXingJieFengAdvancedLogEventData struct {
	beforeLev  int32
	reason     commonlog.DianXingLogReason
	reasonText string
}

func CreatePlayerDianXingJieFengAdvancedLogEventData(beforeLev int32, reason commonlog.DianXingLogReason, reasonText string) *PlayerDianXingJieFengAdvancedLogEventData {
	d := &PlayerDianXingJieFengAdvancedLogEventData{
		beforeLev:  beforeLev,
		reason:     reason,
		reasonText: reasonText,
	}
	return d
}

func (d *PlayerDianXingJieFengAdvancedLogEventData) GetBeforeLev() int32 {
	return d.beforeLev
}

func (d *PlayerDianXingJieFengAdvancedLogEventData) GetReason() commonlog.DianXingLogReason {
	return d.reason
}

func (d *PlayerDianXingJieFengAdvancedLogEventData) GetReasonText() string {
	return d.reasonText
}

//点星系统使用物品
type PlayerDianXingUseItemData struct {
	itemId  int32
	itemNum int32
}

func CreatePlayerDianXingUseItemData(itemId, itemNum int32) *PlayerDianXingUseItemData {
	d := &PlayerDianXingUseItemData{
		itemId:  itemId,
		itemNum: itemNum,
	}
	return d
}

func (d *PlayerDianXingUseItemData) GetItemMap() map[int32]int32 {
	m := map[int32]int32{
		d.itemId: d.itemNum,
	}
	return m
}
