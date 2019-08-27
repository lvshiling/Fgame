package battle

import (
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	worldbosstypes "fgame/fgame/game/worldboss/types"
)

type PlayerBossReliveManager struct {
	p                 scene.Player
	bossReliveDataMap map[worldbosstypes.BossType]*scene.PlayerBossReliveData
}

func (m *PlayerBossReliveManager) GetBossReliveTime(bossType worldbosstypes.BossType) int32 {
	data, ok := m.bossReliveDataMap[bossType]
	if !ok {
		return 0
	}
	return data.GetReliveTime()
}

func (m *PlayerBossReliveManager) PlayerBossRelive(bossType worldbosstypes.BossType) {
	data, ok := m.bossReliveDataMap[bossType]
	if !ok {
		data = scene.CreatePlayerBossReliveData(bossType, 0)
		m.bossReliveDataMap[bossType] = data
	}
	data.Relive()
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerBossReliveDataSync, m.p, data)
}

func (m *PlayerBossReliveManager) PlayerBossReliveSync(bossType worldbosstypes.BossType, reliveTime int32) {
	data, ok := m.bossReliveDataMap[bossType]
	if !ok {
		data = scene.CreatePlayerBossReliveData(bossType, 0)
		m.bossReliveDataMap[bossType] = data
	}
	data.Sync(reliveTime)
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerBossReliveDataSync, m.p, data)
}

func (m *PlayerBossReliveManager) PlayerBossReset(bossType worldbosstypes.BossType) {
	data, ok := m.bossReliveDataMap[bossType]
	if !ok {
		return
	}
	data.Reset()
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerBossReliveDataSync, m.p, data)

}
func (m *PlayerBossReliveManager) GetPlayerBossReliveMap() map[worldbosstypes.BossType]*scene.PlayerBossReliveData {
	return m.bossReliveDataMap

}

func CreatePlayerBossReliveManager(p scene.Player, reliveDataList []*scene.PlayerBossReliveData) *PlayerBossReliveManager {
	m := &PlayerBossReliveManager{
		p: p,
	}
	m.bossReliveDataMap = make(map[worldbosstypes.BossType]*scene.PlayerBossReliveData)
	for _, reliveData := range reliveDataList {
		m.bossReliveDataMap[reliveData.GetBossType()] = reliveData
	}
	return m
}
