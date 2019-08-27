package battle

import (
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//竞技场pvp
type PlayerArenapvpObject struct {
	reliveTimes int32 //已复活次数
}

func CreatePlayerArenapvpObject(reliveTimes int32) *PlayerArenapvpObject {
	obj := &PlayerArenapvpObject{}
	obj.reliveTimes = reliveTimes
	return obj
}

type PlayerArenapvpManager struct {
	p              scene.Player
	arenapvpObj    *PlayerArenapvpObject
	arenapvpBattle bool
}

func (m *PlayerArenapvpManager) GetArenapvpReliveTimes() int32 {
	return m.arenapvpObj.reliveTimes
}

func (m *PlayerArenapvpManager) SetArenapvpReliveTimes(reliveTimes int32) {
	m.arenapvpObj.reliveTimes = reliveTimes
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerArenapvpReliveTimesChanged, m.p, nil)
}

func (m *PlayerArenapvpManager) StartArenapvpBattle() {
	m.arenapvpBattle = true
}

func (m *PlayerArenapvpManager) StopArenapvpBattle() {
	m.arenapvpBattle = false
}

func (m *PlayerArenapvpManager) IsArenapvpBattle() bool {
	return m.arenapvpBattle
}

func CreatePlayerArenapvpManagerWithObject(p scene.Player, arenapvpObj *PlayerArenapvpObject) *PlayerArenapvpManager {
	m := &PlayerArenapvpManager{
		p:           p,
		arenapvpObj: arenapvpObj,
	}
	return m
}

func CreatePlayerArenapvpManager(p scene.Player) *PlayerArenapvpManager {
	m := &PlayerArenapvpManager{
		p: p,
	}
	m.arenapvpObj = CreatePlayerArenapvpObject(0)
	return m
}
