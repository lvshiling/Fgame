package types

import (
	commonlog "fgame/fgame/common/log"
)

type MassacreEventType string

const (
	EventTypeMassacreAdvanced   MassacreEventType = "MassacreAdvanced" //戮仙刃进阶事件
	EventTypeMassacreDegrade                      = "MassacreDegrade"
	EventTypeMassacreChangedLog                   = "MassacreChangedLog" //戮仙刃进阶日志
	EventTypeMassacreWeapon                       = "MassacreWeapon"     //戮仙刃兵魂事件
)

//日志事件类型
type PlayerMassacreChangedLogEventData struct {
	beforeAdvancedNum int32
	changedNum        int32
	beforeShaQiNum    int64
	reason            commonlog.MassacreLogReason
	reasonText        string
}

//兵魂事件类型
type PlayerMassacreWeaponEventData struct {
	weaponId int32
	action   bool
}

func CreatePlayerMassacreChangedLogEventData(beforeAdvancedNum, changedNum int32, beforeShaQiNum int64, reason commonlog.MassacreLogReason, reasonText string) *PlayerMassacreChangedLogEventData {
	d := &PlayerMassacreChangedLogEventData{
		beforeAdvancedNum: beforeAdvancedNum,
		changedNum:        changedNum,
		beforeShaQiNum:    beforeShaQiNum,
		reason:            reason,
		reasonText:        reasonText,
	}
	return d
}

func CreatePlayerMassacreWeaponEventData(weaponId int32, action bool) *PlayerMassacreWeaponEventData {
	d := &PlayerMassacreWeaponEventData{
		weaponId: weaponId,
		action:   action,
	}
	return d
}

func (d *PlayerMassacreChangedLogEventData) GetBeforeAdvancedNum() int32 {
	return d.beforeAdvancedNum
}

func (d *PlayerMassacreChangedLogEventData) GetChangedNum() int32 {
	return d.changedNum
}

func (d *PlayerMassacreChangedLogEventData) GetBeforeShaQiNum() int64 {
	return d.beforeShaQiNum
}

func (d *PlayerMassacreChangedLogEventData) GetReason() commonlog.MassacreLogReason {
	return d.reason
}

func (d *PlayerMassacreChangedLogEventData) GetReasonText() string {
	return d.reasonText
}

func (d *PlayerMassacreWeaponEventData) GetWeaponId() int32 {
	return d.weaponId
}

func (d *PlayerMassacreWeaponEventData) GetAction() bool {
	return d.action
}

type PlayerMassacreDegradeEventData struct {
	oldAdvanceId int32
	newAdvanceId int32
	attackName   string
}

func (d *PlayerMassacreDegradeEventData) GetOldAdvanceId() int32 {
	return d.oldAdvanceId
}

func (d *PlayerMassacreDegradeEventData) GetNewAdvanceId() int32 {
	return d.newAdvanceId
}

func (d *PlayerMassacreDegradeEventData) GetAttackName() string {
	return d.attackName
}

func CreatePlayerMassacreDegradeEventData(oldAdvanceId int32, newAdvanceId int32, attackName string) *PlayerMassacreDegradeEventData {
	d := &PlayerMassacreDegradeEventData{
		oldAdvanceId: oldAdvanceId,
		newAdvanceId: newAdvanceId,
		attackName:   attackName,
	}
	return d
}

type PlayerMassacreAdvancedEventData struct {
	oldAdvanceId int32
	newAdvanceId int32
	attackName   string
}

func (d *PlayerMassacreAdvancedEventData) GetOldAdvanceId() int32 {
	return d.oldAdvanceId
}

func (d *PlayerMassacreAdvancedEventData) GetNewAdvanceId() int32 {
	return d.newAdvanceId
}

func CreatePlayerMassacreAdvanceEventData(oldAdvanceId int32, newAdvanceId int32) *PlayerMassacreAdvancedEventData {
	d := &PlayerMassacreAdvancedEventData{
		oldAdvanceId: oldAdvanceId,
		newAdvanceId: newAdvanceId,
	}
	return d
}
