package types

import (
	commonlogtypes "fgame/fgame/common/log/types"
	propertytypes "fgame/fgame/game/property/types"
)

type PropertyEventType string

const (
	EventTypePlayerExpAdd                  PropertyEventType = "PlayerExpAdd"
	EventTypePlayerPropertyEffectorChanged                   = "PlayerPropertyEffectorChanged"
	EventTypePlayerLevelChanged                              = "PlayerLevelChanged"
	EventTypePlayerZhuanShengChanged                         = "PlayerZhuanShengChanged"
	EventTypePlayerSystemPropertyChanged                     = "PlayerSystemPropertyChanged"
	EventTypePlayerGoldCost                                  = "PlayerGoldCost"
	EventTypePlayerGoldChangedLog                            = "PlayerGoldChangedLog"
	EventTypePlayerNewGoldChangedLog                         = "PlayerNewGoldChangedLog"
	EventTypePlayerNewBindGoldChangedLog                     = "PlayerNewBindGoldChangedLog"
	EventTypePlayerSilverChangedLog                          = "PlayerSilverChangedLog"
	EventTypePlayerCharmAdd                                  = "PlayerCharmAdd"
	EventTypePlayerForceChanged                              = "PlayerForceChanged"
	EventTypePlayerExpChangedLog                             = "PlayerExpChangedLog"
	EventTypePlayerGoldCostIncludeBind                       = "PlayerGoldCostIncludeBind"
)

func CreatePropertyChangedEventData(propertyType propertytypes.BasePropertyType, val int64) *BasePropertyChangedEventData {
	bpced := &BasePropertyChangedEventData{

		propertyType:  propertyType,
		propertyValue: val,
	}
	return bpced
}

type BasePropertyChangedEventData struct {
	propertyType  propertytypes.BasePropertyType
	propertyValue int64
}

func (qced *BasePropertyChangedEventData) PropertyType() propertytypes.BasePropertyType {
	return qced.propertyType
}

func (qced *BasePropertyChangedEventData) PropertyValue() int64 {
	return qced.propertyValue
}

//元宝变化
type PlayerGoldChangedLogEventData struct {
	beforeGold     int64
	beforeBindGold int64
	changedNum     int64
	reason         commonlogtypes.LogReason
	reasonText     string
}

func CreatePlayerGoldChangedLogEventData(beforeGold, beforeBindGold, changedNum int64, reason commonlogtypes.LogReason, reasonText string) *PlayerGoldChangedLogEventData {
	d := &PlayerGoldChangedLogEventData{
		beforeGold:     beforeGold,
		beforeBindGold: beforeBindGold,
		changedNum:     changedNum,
		reason:         reason,
		reasonText:     reasonText,
	}
	return d
}

func (d *PlayerGoldChangedLogEventData) GetBeforeGold() int64 {
	return d.beforeGold
}

func (d *PlayerGoldChangedLogEventData) GetBeforeBindGold() int64 {
	return d.beforeBindGold
}

func (d *PlayerGoldChangedLogEventData) GetChangedNum() int64 {
	return d.changedNum
}

func (d *PlayerGoldChangedLogEventData) GetReason() commonlogtypes.LogReason {
	return d.reason
}

func (d *PlayerGoldChangedLogEventData) GetReasonText() string {
	return d.reasonText
}

//
type PlayerSilverChangedLogEventData struct {
	beforeSilver int64
	changedNum   int64
	reason       commonlogtypes.LogReason
	reasonText   string
}

func CreatePlayerSilverChangedLogEventData(beforeSilver, changedNum int64, reason commonlogtypes.LogReason, reasonText string) *PlayerSilverChangedLogEventData {
	d := &PlayerSilverChangedLogEventData{
		beforeSilver: beforeSilver,
		changedNum:   changedNum,
		reason:       reason,
		reasonText:   reasonText,
	}
	return d
}

func (d *PlayerSilverChangedLogEventData) GetBeforeSilver() int64 {
	return d.beforeSilver
}

func (d *PlayerSilverChangedLogEventData) GetChangedNum() int64 {
	return d.changedNum
}

func (d *PlayerSilverChangedLogEventData) GetReason() commonlogtypes.LogReason {
	return d.reason
}

func (d *PlayerSilverChangedLogEventData) GetReasonText() string {
	return d.reasonText
}

//新元宝变化
type PlayerNewGoldChangedLogEventData struct {
	beforeGold int64
	changedNum int64
	reason     commonlogtypes.LogReason
	reasonText string
}

func CreatePlayerNewGoldChangedLogEventData(beforeGold, changedNum int64, reason commonlogtypes.LogReason, reasonText string) *PlayerNewGoldChangedLogEventData {
	d := &PlayerNewGoldChangedLogEventData{
		beforeGold: beforeGold,
		changedNum: changedNum,
		reason:     reason,
		reasonText: reasonText,
	}
	return d
}

func (d *PlayerNewGoldChangedLogEventData) GetBeforeGold() int64 {
	return d.beforeGold
}

func (d *PlayerNewGoldChangedLogEventData) GetChangedNum() int64 {
	return d.changedNum
}

func (d *PlayerNewGoldChangedLogEventData) GetReason() commonlogtypes.LogReason {
	return d.reason
}

func (d *PlayerNewGoldChangedLogEventData) GetReasonText() string {
	return d.reasonText
}

//绑元变化变化
type PlayerNewBindGoldChangedLogEventData struct {
	beforeBindGold int64
	changedNum     int64
	reason         commonlogtypes.LogReason
	reasonText     string
}

func CreatePlayerNewBindGoldChangedLogEventData(beforeBindGold, changedNum int64, reason commonlogtypes.LogReason, reasonText string) *PlayerNewBindGoldChangedLogEventData {
	d := &PlayerNewBindGoldChangedLogEventData{
		beforeBindGold: beforeBindGold,
		changedNum:     changedNum,
		reason:         reason,
		reasonText:     reasonText,
	}
	return d
}

func (d *PlayerNewBindGoldChangedLogEventData) GetBeforeBindGold() int64 {
	return d.beforeBindGold
}

func (d *PlayerNewBindGoldChangedLogEventData) GetChangedNum() int64 {
	return d.changedNum
}

func (d *PlayerNewBindGoldChangedLogEventData) GetReason() commonlogtypes.LogReason {
	return d.reason
}

func (d *PlayerNewBindGoldChangedLogEventData) GetReasonText() string {
	return d.reasonText
}

//战力变化变化
type PlayerForceChangedEventData struct {
	beforeForce int64
	force       int64
	mask        uint64
	reasonText  string
}

func CreatePlayerForceChangedEventData(beforeForce, force int64, mask uint64, reasonText string) *PlayerForceChangedEventData {
	d := &PlayerForceChangedEventData{
		beforeForce: beforeForce,
		force:       force,
		mask:        mask,
		reasonText:  reasonText,
	}
	return d
}

func (d *PlayerForceChangedEventData) GetBeforeForce() int64 {
	return d.beforeForce
}

func (d *PlayerForceChangedEventData) GetForce() int64 {
	return d.force
}

func (d *PlayerForceChangedEventData) GetMask() uint64 {
	return d.mask
}

func (d *PlayerForceChangedEventData) GetReasonText() string {
	return d.reasonText
}

type PlayerExpChangedLogEventData struct {
	beforeExp   int64
	curExp      int64
	beforeLevel int32
	curLevel    int32
	changedExp  int64
	reason      commonlogtypes.LogReason
	reasonText  string
}

func CreatePlayerExpChangedLogEventData(beforeExp, curExp, changedExp int64, beforeLevel, curLevel int32, reason commonlogtypes.LogReason, reasonText string) *PlayerExpChangedLogEventData {
	d := &PlayerExpChangedLogEventData{
		beforeExp:   beforeExp,
		curExp:      curExp,
		beforeLevel: beforeLevel,
		curLevel:    curLevel,
		changedExp:  changedExp,
		reason:      reason,
		reasonText:  reasonText,
	}
	return d
}

func (d *PlayerExpChangedLogEventData) GetBeforeExp() int64 {
	return d.beforeExp
}

func (d *PlayerExpChangedLogEventData) GetCurExp() int64 {
	return d.curExp
}

func (d *PlayerExpChangedLogEventData) GetBeforeLevel() int32 {
	return d.beforeLevel
}

func (d *PlayerExpChangedLogEventData) GetCurLevel() int32 {
	return d.curLevel
}

func (d *PlayerExpChangedLogEventData) GetChangedExp() int64 {
	return d.changedExp
}

func (d *PlayerExpChangedLogEventData) GetReason() commonlogtypes.LogReason {
	return d.reason
}

func (d *PlayerExpChangedLogEventData) GetReasonText() string {
	return d.reasonText
}
