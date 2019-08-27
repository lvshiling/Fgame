package battle

import (
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
)

// 幻境boss
type PlayerUnrealBossManager struct {
	p                    scene.Player
	pilaoNum             int32
	lastAttackNoticeTime int64
}

const (
	cd = int64(10 * common.SECOND)
)

func (m *PlayerUnrealBossManager) IsEnoughPilao(pilao int32) bool {
	return m.pilaoNum >= pilao
}
func (m *PlayerUnrealBossManager) GetPilao() int32 {
	return m.pilaoNum
}

func (m *PlayerUnrealBossManager) IsPilaoNoticeCd() bool {
	now := global.GetGame().GetTimeService().Now()
	if now-m.lastAttackNoticeTime <= cd {
		return true
	}

	m.lastAttackNoticeTime = now
	return false
}

func (m *PlayerUnrealBossManager) SynPilaoNum(pilao int32) {
	m.pilaoNum = pilao
}

func CreatePlayerUnrealBossManagerWithData(p scene.Player, pilao int32) *PlayerUnrealBossManager {
	m := &PlayerUnrealBossManager{
		p:        p,
		pilaoNum: pilao,
	}
	return m
}

func CreatePlayerUnrealBossManager(p scene.Player) *PlayerUnrealBossManager {
	m := &PlayerUnrealBossManager{
		p: p,
	}
	return m
}
