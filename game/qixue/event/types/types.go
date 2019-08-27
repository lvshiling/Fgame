package types

import (
	commonlog "fgame/fgame/common/log"
)

type QiXueEventType string

const (
	EventTypeQiXueAdvanced QiXueEventType = "QiXueAdvanced" //泣血枪进阶事件
	EventTypeQiXueDegrade                 = "QiXueDegrade"
)

// 日志事件
type QuXueLogEventType string

const (
	EventTypeQiXueChangedLog QuXueLogEventType = "QiXueChangedLog" //泣血枪进阶日志
)

//日志事件类型
type PlayerQiXueChangedLogEventData struct {
	beforeAdvancedNum int32
	changedNum        int32
	beforeShaLuNum    int64
	reason            commonlog.QiXueLogReason
	reasonText        string
}

func CreatePlayerQiXueChangedLogEventData(beforeAdvancedNum, changedNum int32, beforeShaLuNum int64, reason commonlog.QiXueLogReason, reasonText string) *PlayerQiXueChangedLogEventData {
	d := &PlayerQiXueChangedLogEventData{
		beforeAdvancedNum: beforeAdvancedNum,
		changedNum:        changedNum,
		beforeShaLuNum:    beforeShaLuNum,
		reason:            reason,
		reasonText:        reasonText,
	}
	return d
}

func (d *PlayerQiXueChangedLogEventData) GetBeforeAdvancedNum() int32 {
	return d.beforeAdvancedNum
}

func (d *PlayerQiXueChangedLogEventData) GetChangedNum() int32 {
	return d.changedNum
}

func (d *PlayerQiXueChangedLogEventData) GetBeforeShaLuNum() int64 {
	return d.beforeShaLuNum
}

func (d *PlayerQiXueChangedLogEventData) GetReason() commonlog.QiXueLogReason {
	return d.reason
}

func (d *PlayerQiXueChangedLogEventData) GetReasonText() string {
	return d.reasonText
}

//泣血阶数变化
type PlayerQiXueChangedWeaponEventData struct {
	oldLevel int32
	oldStar  int32
	newLevel int32
	newStar  int32
}

func (d *PlayerQiXueChangedWeaponEventData) GetOldLevel() int32 {
	return d.oldLevel
}

func (d *PlayerQiXueChangedWeaponEventData) GetOldStar() int32 {
	return d.oldStar
}

func (d *PlayerQiXueChangedWeaponEventData) GetNewLevel() int32 {
	return d.newLevel
}

func (d *PlayerQiXueChangedWeaponEventData) GetNewStar() int32 {
	return d.newStar
}

func CreatePlayerQiXueChangedWeaponEventData(oldLevel int32, oldStar int32, newLevel int32, newStar int32) *PlayerQiXueChangedWeaponEventData {
	d := &PlayerQiXueChangedWeaponEventData{
		oldLevel: oldLevel,
		oldStar:  oldStar,
		newLevel: newLevel,
		newStar:  newStar,
	}
	return d
}

//兵魂事件类型
type PlayerQiXueWeaponEventData struct {
	weaponId int32
	action   bool
}

func (d *PlayerQiXueWeaponEventData) GetWeaponId() int32 {
	return d.weaponId
}

func (d *PlayerQiXueWeaponEventData) GetAction() bool {
	return d.action
}

func CreatePlayerQiXueWeaponEventData(weaponId int32, action bool) *PlayerQiXueWeaponEventData {
	d := &PlayerQiXueWeaponEventData{
		weaponId: weaponId,
		action:   action,
	}
	return d
}
