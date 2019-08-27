package battle

import (
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

type PlayerDenseWatManager struct {
	p       scene.Player
	num     int32
	endTime int64
}

func (m *PlayerDenseWatManager) GetDenseWatNum() int32 {
	return m.num
}

func (m *PlayerDenseWatManager) SetDenseWatNum(num int32) {
	if num < 0 {
		return
	}
	m.num = num
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerDenseWatNumChanged, m.p, nil)
}

func (m *PlayerDenseWatManager) GetDenseWatEndTime() int64 {
	return m.endTime
}

func (m *PlayerDenseWatManager) SetDenseWatEndTime(endTime int64) {
	if endTime <= 0 {
		return
	}
	m.endTime = endTime
	m.num = 0
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerDenseWatEndTimeSet, m.p, nil)
}

func (m *PlayerDenseWatManager) SyncDenseWat(num int32, endTime int64) {
	m.num = num
	m.endTime = endTime
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerDenseWatSync, m.p, nil)
}

func CreatePlayerDenseWatManager(p scene.Player, num int32, endTime int64) *PlayerDenseWatManager {
	m := &PlayerDenseWatManager{}
	m.p = p
	m.num = num
	m.endTime = endTime
	return m
}
