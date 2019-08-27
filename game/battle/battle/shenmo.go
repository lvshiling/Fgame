package battle

import (
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

type PlayerShenMoManager struct {
	p          scene.Player
	gongXunNum int32
	killNum    int32
	endTime    int64
}

func (m *PlayerShenMoManager) GetShenMoGongXunNum() int32 {
	return m.gongXunNum
}

func (m *PlayerShenMoManager) SetShenMoGongXunNum(gongXunNum int32) {
	if gongXunNum < 0 {
		return
	}
	m.gongXunNum = gongXunNum
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShenMoGongXunNumChanged, m.p, nil)
}

func (m *PlayerShenMoManager) GetShenMoKillNum() int32 {
	return m.killNum
}

func (m *PlayerShenMoManager) SetShenMoKillNum(killNum int32) {
	if killNum < 0 {
		return
	}
	m.killNum = killNum
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShenMoKillNumChanged, m.p, nil)
}

func (m *PlayerShenMoManager) GetShenMoEndTime() int64 {
	return m.endTime
}

func (m *PlayerShenMoManager) SetShenMoEndTime(endTime int64) {
	if endTime <= 0 {
		return
	}
	m.endTime = endTime
	m.killNum = 0
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShenMoEndTimeSet, m.p, nil)
}

func (m *PlayerShenMoManager) SyncShenMo(gongXunNum int32, killNum int32, endTime int64) {
	m.gongXunNum = gongXunNum
	m.killNum = killNum
	m.endTime = endTime
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerShenMoSync, m.p, nil)
}

func CreatePlayerShenMoManager(p scene.Player, gongXunNum int32, killNum int32, endTime int64) *PlayerShenMoManager {
	m := &PlayerShenMoManager{}
	m.p = p
	m.gongXunNum = gongXunNum
	m.killNum = killNum
	m.endTime = endTime
	return m
}

func CreatePlayerShenMoManagerBase(p scene.Player) *PlayerShenMoManager {
	m := &PlayerShenMoManager{}
	m.p = p

	return m
}
