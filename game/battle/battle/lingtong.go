package battle

import (
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

type LingTongChangeEventData struct {
	oldLingTong scene.LingTong
	newLingTong scene.LingTong
}

func (d *LingTongChangeEventData) GetOldLingTong() scene.LingTong {
	return d.oldLingTong
}

func (d *LingTongChangeEventData) GetNewLingTong() scene.LingTong {
	return d.newLingTong
}

func CreateLingTongChangeEventData(oldLingTong, newLingTong scene.LingTong) *LingTongChangeEventData {
	d := &LingTongChangeEventData{
		oldLingTong: oldLingTong,
		newLingTong: newLingTong,
	}
	return d
}

type PlayerLingTongShowManager struct {
	p        scene.Player
	lingTong scene.LingTong
	hidden   bool
}

func (m *PlayerLingTongShowManager) UpdateLingTong(lingTong scene.LingTong) {
	oldLingTong := m.lingTong
	m.lingTong = lingTong
	d := CreateLingTongChangeEventData(oldLingTong, lingTong)
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerLingTongChanged, m.p, d)
}

func (m *PlayerLingTongShowManager) HiddenLingTong(flag bool) {
	if flag == m.hidden {
		return
	}
	m.hidden = flag
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerLingTongShow, m.p, nil)
}

func (m *PlayerLingTongShowManager) IsLingTongHidden() bool {
	return m.hidden
}

func (m *PlayerLingTongShowManager) GetLingTong() scene.LingTong {
	return m.lingTong
}

func CreatePlayerLingTongShowManagerWithLingTong(p scene.Player, lingTong scene.LingTong) *PlayerLingTongShowManager {
	m := &PlayerLingTongShowManager{
		p:        p,
		lingTong: lingTong,
	}
	return m
}

func CreatePlayerLingTongShowManager(p scene.Player) *PlayerLingTongShowManager {
	m := &PlayerLingTongShowManager{
		p: p,
	}
	return m
}
