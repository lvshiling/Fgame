package battle

import (
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

type PlayerCollectManager struct {
	p scene.Player
	n scene.CollectNPC
}

func (m *PlayerCollectManager) Collect(n scene.CollectNPC) (flag bool) {
	if m.n != nil {
		return
	}
	m.n = n
	flag = true
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerCollect, m.p, nil)
	return
}

func (m *PlayerCollectManager) HasCollect() (n scene.CollectNPC, flag bool) {
	if m.n == nil {
		return
	}
	n = m.n
	flag = true
	return
}

func (m *PlayerCollectManager) ClearCollect() {
	if m.n == nil {
		return
	}
	m.n = nil
}

func CreatePlayerCollectManager(p scene.Player) *PlayerCollectManager {
	m := &PlayerCollectManager{}
	m.p = p
	return m
}
