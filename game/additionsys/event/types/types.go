package types

import (
	commonlog "fgame/fgame/common/log"
	additionsystypes "fgame/fgame/game/additionsys/types"
)

type AdditionSysEventType string

const (
	EventTypeAdditionSysStrengthenLevLog AdditionSysEventType = "AdditionSysStrengthenLevChangeLog" // 强化等级日志
	EventTypeAdditionSysShengJi                               = "AdditionSysShengJi"                // 系统等级升级
	EventTypeAdditionSysShengJiLog                            = "AdditionSysShengJiLog"             // 系统等级升级日志
	EventTypeAdditionSysUpgrade                               = "AdditionUpgrade"                   // 系统进阶
	EventTypeAdditionSysShenZhuLog                            = "AdditionSysShenZhuLog"             // 系统神铸升级日志
	EventTypeAdditionSysTongLingLog                           = "AdditionSysTongLingLog"            // 系统通灵升级日志
	EventTypeAdditionSysAwakeLog                              = "AdditionSysAwakeLog"               // 系统觉醒日志
	EventTypeAdditionSysUseItem                               = "AdditionSysUseItem"                // 系统消耗道具(升级、进阶等等)
)

//强化等级日志事件类型
type PlayerAdditionSysStrengthenLevLogEventData struct {
	sysType    additionsystypes.AdditionSysType
	position   additionsystypes.SlotPositionType
	beforeLev  int32
	reason     commonlog.AdditionSysLogReason
	reasonText string
}

func CreatePlayerAdditionSysStrengthenLevLogEventData(sysType additionsystypes.AdditionSysType, position additionsystypes.SlotPositionType, beforeLev int32, reason commonlog.AdditionSysLogReason, reasonText string) *PlayerAdditionSysStrengthenLevLogEventData {
	d := &PlayerAdditionSysStrengthenLevLogEventData{
		sysType:    sysType,
		position:   position,
		beforeLev:  beforeLev,
		reason:     reason,
		reasonText: reasonText,
	}
	return d
}

func (d *PlayerAdditionSysStrengthenLevLogEventData) GetSysType() additionsystypes.AdditionSysType {
	return d.sysType
}

func (d *PlayerAdditionSysStrengthenLevLogEventData) GetPosition() additionsystypes.SlotPositionType {
	return d.position
}

func (d *PlayerAdditionSysStrengthenLevLogEventData) GetBeforeLev() int32 {
	return d.beforeLev
}

func (d *PlayerAdditionSysStrengthenLevLogEventData) GetReason() commonlog.AdditionSysLogReason {
	return d.reason
}

func (d *PlayerAdditionSysStrengthenLevLogEventData) GetReasonText() string {
	return d.reasonText
}

//系统等级升级日志
type PlayerAdditionSysShengJiLogEventData struct {
	sysType     additionsystypes.AdditionSysType
	beforeLev   int32
	beforeUpNum int32
	beforeUpPro int32
	reason      commonlog.AdditionSysLogReason
	reasonText  string
}

func CreatePlayerAdditionSysShengJiLogEventData(sysType additionsystypes.AdditionSysType, beforeLev int32, beforeUpNum int32, beforeUpPro int32, reason commonlog.AdditionSysLogReason, reasonText string) *PlayerAdditionSysShengJiLogEventData {
	d := &PlayerAdditionSysShengJiLogEventData{
		sysType:     sysType,
		beforeLev:   beforeLev,
		beforeUpNum: beforeUpNum,
		beforeUpPro: beforeUpPro,
		reason:      reason,
		reasonText:  reasonText,
	}
	return d
}

func (d *PlayerAdditionSysShengJiLogEventData) GetSysType() additionsystypes.AdditionSysType {
	return d.sysType
}

func (d *PlayerAdditionSysShengJiLogEventData) GetBeforeLev() int32 {
	return d.beforeLev
}

func (d *PlayerAdditionSysShengJiLogEventData) GetBeforeUpNum() int32 {
	return d.beforeUpNum
}

func (d *PlayerAdditionSysShengJiLogEventData) GetBeforeUpPro() int32 {
	return d.beforeUpPro
}

func (d *PlayerAdditionSysShengJiLogEventData) GetReason() commonlog.AdditionSysLogReason {
	return d.reason
}

func (d *PlayerAdditionSysShengJiLogEventData) GetReasonText() string {
	return d.reasonText
}

//系统进阶日志
type PlayerAdditionUpgradeEventData struct {
	sysType additionsystypes.AdditionSysType
	pos     additionsystypes.SlotPositionType
}

func CreatePlayerAdditionUpgradeEventData(sysType additionsystypes.AdditionSysType, pos additionsystypes.SlotPositionType) *PlayerAdditionUpgradeEventData {
	d := &PlayerAdditionUpgradeEventData{
		sysType: sysType,
		pos:     pos,
	}
	return d
}

func (d *PlayerAdditionUpgradeEventData) GetSysType() additionsystypes.AdditionSysType {
	return d.sysType
}

func (d *PlayerAdditionUpgradeEventData) GetPos() additionsystypes.SlotPositionType {
	return d.pos
}

//神铸升级日志事件类型
type PlayerAdditionSysShenZhuLevLogEventData struct {
	sysType    additionsystypes.AdditionSysType
	position   additionsystypes.SlotPositionType
	beforeLev  int32
	reason     commonlog.AdditionSysLogReason
	reasonText string
}

func CreatePlayerAdditionSysShenZhuLevLogEventData(typ additionsystypes.AdditionSysType, pos additionsystypes.SlotPositionType, beforeLev int32, reason commonlog.AdditionSysLogReason, reasonText string) *PlayerAdditionSysShenZhuLevLogEventData {
	d := &PlayerAdditionSysShenZhuLevLogEventData{
		sysType:    typ,
		position:   pos,
		beforeLev:  beforeLev,
		reason:     reason,
		reasonText: reasonText,
	}
	return d
}

func (d *PlayerAdditionSysShenZhuLevLogEventData) GetSysType() additionsystypes.AdditionSysType {
	return d.sysType
}

func (d *PlayerAdditionSysShenZhuLevLogEventData) GetPosition() additionsystypes.SlotPositionType {
	return d.position
}

func (d *PlayerAdditionSysShenZhuLevLogEventData) GetBeforeLev() int32 {
	return d.beforeLev
}

func (d *PlayerAdditionSysShenZhuLevLogEventData) GetReason() commonlog.AdditionSysLogReason {
	return d.reason
}

func (d *PlayerAdditionSysShenZhuLevLogEventData) GetReasonText() string {
	return d.reasonText
}

//系统通灵升级日志
type PlayerAdditionSysTongLingLogEventData struct {
	sysType    additionsystypes.AdditionSysType
	beforeLev  int32
	reason     commonlog.AdditionSysLogReason
	reasonText string
}

func CreatePlayerAdditionSysTongLingLogEventData(sysType additionsystypes.AdditionSysType, beforeLev int32, reason commonlog.AdditionSysLogReason, reasonText string) *PlayerAdditionSysTongLingLogEventData {
	d := &PlayerAdditionSysTongLingLogEventData{
		sysType:    sysType,
		beforeLev:  beforeLev,
		reason:     reason,
		reasonText: reasonText,
	}
	return d
}

func (d *PlayerAdditionSysTongLingLogEventData) GetSysType() additionsystypes.AdditionSysType {
	return d.sysType
}

func (d *PlayerAdditionSysTongLingLogEventData) GetBeforeLev() int32 {
	return d.beforeLev
}

func (d *PlayerAdditionSysTongLingLogEventData) GetReason() commonlog.AdditionSysLogReason {
	return d.reason
}

func (d *PlayerAdditionSysTongLingLogEventData) GetReasonText() string {
	return d.reasonText
}

//系统觉醒日志
type PlayerAdditionSysAwakeEventData struct {
	sysType       additionsystypes.AdditionSysType
	beforeIsAwake int32
	reason        commonlog.AdditionSysLogReason
	reasonText    string
}

func CreatePlayerAdditionSysAwakeLogEventData(sysType additionsystypes.AdditionSysType, isAwake int32, reason commonlog.AdditionSysLogReason, reasonText string) *PlayerAdditionSysAwakeEventData {
	d := &PlayerAdditionSysAwakeEventData{
		sysType:       sysType,
		beforeIsAwake: isAwake,
		reason:        reason,
		reasonText:    reasonText,
	}
	return d
}

func (d *PlayerAdditionSysAwakeEventData) GetSysType() additionsystypes.AdditionSysType {
	return d.sysType
}

func (d *PlayerAdditionSysAwakeEventData) GetBeforeIsAwake() int32 {
	return d.beforeIsAwake
}

func (d *PlayerAdditionSysAwakeEventData) GetReason() commonlog.AdditionSysLogReason {
	return d.reason
}

func (d *PlayerAdditionSysAwakeEventData) GetReasonText() string {
	return d.reasonText
}

//系统使用道具
type PlayerAdditionSysUseItemEventData struct {
	itemMap map[int32]int32
}

func CreatePlayerAdditionSysUseItemEventData(itemMap map[int32]int32) *PlayerAdditionSysUseItemEventData {
	d := &PlayerAdditionSysUseItemEventData{
		itemMap: itemMap,
	}
	return d
}

func (data *PlayerAdditionSysUseItemEventData) GetItemMap() map[int32]int32 {
	return data.itemMap
}
