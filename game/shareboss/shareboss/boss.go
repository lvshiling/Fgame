package shareboss

import (
	coretypes "fgame/fgame/core/types"
)

type ShareBossInfo struct {
	pos       coretypes.Position
	biologyId int32
	isDead    bool
	deadTime  int64
}

func (info *ShareBossInfo) GetPosition() coretypes.Position {
	return info.pos
}

func (info *ShareBossInfo) GetBiologyId() int32 {
	return info.biologyId
}

func (info *ShareBossInfo) IsDead() bool {
	return info.isDead
}

func (info *ShareBossInfo) GetDeadTime() int64 {
	return info.deadTime
}
