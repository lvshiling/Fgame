package battle

import (
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//竞技场
type PlayerArenaObject struct {
	reliveTime int32
	winTime    int32
}

func CreatePlayerArenaObject(
	reliveTime int32,
	winTime int32,
) *PlayerArenaObject {
	obj := &PlayerArenaObject{}
	obj.reliveTime = reliveTime
	obj.winTime = winTime
	return obj
}

type PlayerArenaManager struct {
	p           scene.Player
	arenaObj    *PlayerArenaObject
	arenaBattle bool
}

func (m *PlayerArenaManager) GetArenaReliveTime() int32 {
	return m.arenaObj.reliveTime
}

func (m *PlayerArenaManager) SetArenaReliveTime(reliveTime int32) {
	m.arenaObj.reliveTime = reliveTime
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerArenaReliveTimeChanged, m.p, nil)
}

func (m *PlayerArenaManager) GetArenaWinTime() int32 {
	return m.arenaObj.winTime
}

func (m *PlayerArenaManager) SetArenaWinTime(winTime int32) {
	m.arenaObj.winTime = winTime
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerArenaWinTimeChanged, m.p, nil)
}

func (m *PlayerArenaManager) StartArenaBattle() {
	m.arenaBattle = true
}

func (m *PlayerArenaManager) StopArenaBattle() {
	m.arenaBattle = false
}

func (m *PlayerArenaManager) IsArenaBattle() bool {
	return m.arenaBattle
}

func CreatePlayerArenaManagerWithObject(p scene.Player, arenaObj *PlayerArenaObject) *PlayerArenaManager {
	m := &PlayerArenaManager{
		p:        p,
		arenaObj: arenaObj,
	}
	return m
}
