package battle

import (
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
)

//状态数据管理器
type StateDataManager struct {
	bo                scene.BattleObject
	skillTime         int64
	skillActionTime   int64
	skilledTime       int64
	destPos           coretypes.Position
	skilledStopTime   int64
	attackedMoveSpeed float64
}

func (m *StateDataManager) SetSkillActionTime(skillActionTime int64) {
	now := global.GetGame().GetTimeService().Now()
	m.skillTime = now
	m.skillActionTime = skillActionTime
}

func (m *StateDataManager) GetSkillTime() int64 {
	return m.skillTime
}

func (m *StateDataManager) GetSkillActionTime() int64 {
	return m.skillActionTime
}

func (m *StateDataManager) GetDestPosition() coretypes.Position {
	return m.destPos
}

func (m *StateDataManager) SkilledStop(destPos coretypes.Position, skilledStopTime int64, attackedMoveSpeed float64) {
	now := global.GetGame().GetTimeService().Now()
	m.destPos = destPos
	m.skilledTime = now
	m.skilledStopTime = skilledStopTime
	m.attackedMoveSpeed = attackedMoveSpeed
}

func (m *StateDataManager) GetSkilledTime() int64 {
	return m.skilledTime
}

func (m *StateDataManager) GetSkilledStopTime() int64 {
	return m.skilledStopTime
}

func (m *StateDataManager) GetAttackedMoveSpeed() float64 {
	return m.attackedMoveSpeed
}

func CreateStateDateManager(bo scene.BattleObject) *StateDataManager {
	m := &StateDataManager{}
	m.bo = bo
	return m
}
