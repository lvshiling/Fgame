package battle

import (
	godsiegetypes "fgame/fgame/game/godsiege/types"
	"fgame/fgame/game/scene/scene"
)

type PlayerGodSiegeManager struct {
	p       scene.Player
	godType godsiegetypes.GodSiegeType
}

func (m *PlayerGodSiegeManager) IsGodSiegeLineUp() bool {
	return m.godType != godsiegetypes.GodSiegeTypeNo
}

func (m *PlayerGodSiegeManager) GetGodSiegeLineUp() (isLineUp bool, godType godsiegetypes.GodSiegeType) {
	if m.godType == godsiegetypes.GodSiegeTypeNo {
		return
	}
	isLineUp = true
	godType = m.godType
	return
}

func (m *PlayerGodSiegeManager) GodSiegeCancleLineUp() (flag bool) {
	if m.godType == godsiegetypes.GodSiegeTypeNo {
		return
	}
	flag = true
	m.godType = godsiegetypes.GodSiegeTypeNo
	return
}

func (m *PlayerGodSiegeManager) GodSiegeLineUp(godType godsiegetypes.GodSiegeType) (flag bool) {
	if m.godType != godsiegetypes.GodSiegeTypeNo {
		return
	}
	flag = true
	m.godType = godType
	return
}

func CreatePlayerGodSiegeManager(p scene.Player) *PlayerGodSiegeManager {
	m := &PlayerGodSiegeManager{}
	m.p = p
	m.godType = godsiegetypes.GodSiegeTypeNo
	return m
}
