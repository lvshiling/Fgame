package types

import (
	coretypes "fgame/fgame/core/types"
)

const (
	EventTypeBattleObjectSkillTrigger     BattleEventType = "BattleObjectSkillTrigger"
	EventTypeBattleObjectMoveTrigger                      = "BattleObjectMoveTrigger"
	EventTypeBattleObjectHuDun                            = "BattleObjectHuDun"
	EventTypeBattleObjectHateChanged                      = "BattleObjectHateChanged"
	EventTypeBattleObjectMoveFix                          = "BattleObjectMoveFix"
	EventTypeBattleObjectMaxDamageChanged                 = "BattleObjectMaxDamageChanged"
	EventTypeBattleObjectTeshuSkillReset                  = "BattleObjectTeshuSkillReset"
)

type BattleObjectMoveTriggerEventData struct {
	destPos coretypes.Position
	angle   float64
	speed   float64
}

func (d *BattleObjectMoveTriggerEventData) GetAngle() float64 {
	return d.angle
}

func (d *BattleObjectMoveTriggerEventData) GetDestPos() coretypes.Position {
	return d.destPos
}

func (d *BattleObjectMoveTriggerEventData) GetSpeed() float64 {
	return d.speed
}

func CreateBattleObjectMoveTriggerEventData(destPos coretypes.Position, angle float64, speed float64) *BattleObjectMoveTriggerEventData {
	d := &BattleObjectMoveTriggerEventData{}
	d.angle = angle
	d.destPos = destPos
	d.speed = speed
	return d
}

type BattleObjectMaxDamageChangedEventData struct {
	originAttackId  int64
	currentAttackId int64
}

func (d *BattleObjectMaxDamageChangedEventData) GetOriginAttackId() int64 {
	return d.originAttackId
}

func (d *BattleObjectMaxDamageChangedEventData) GetCurrentAttackId() int64 {
	return d.currentAttackId
}

func CreateBattleObjectMaxDamageChangedEventData(originAttackId int64, currentAttackId int64) *BattleObjectMaxDamageChangedEventData {
	d := &BattleObjectMaxDamageChangedEventData{}
	d.originAttackId = originAttackId
	d.currentAttackId = currentAttackId

	return d
}
