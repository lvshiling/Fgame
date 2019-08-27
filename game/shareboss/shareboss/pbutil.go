package shareboss

import (
	coretypes "fgame/fgame/core/types"
	sharebosspb "fgame/fgame/cross/shareboss/pb"
)

func convertFromPos(pos *sharebosspb.Position) coretypes.Position {
	tpos := coretypes.Position{}
	tpos.X = float64(pos.X)
	tpos.Y = float64(pos.Y)
	tpos.Z = float64(pos.Z)
	return tpos
}

func convertFromBossInfo(bossInfo *sharebosspb.BossInfo) *ShareBossInfo {
	shareBossInfo := &ShareBossInfo{}
	shareBossInfo.pos = convertFromPos(bossInfo.GetPos())
	shareBossInfo.biologyId = bossInfo.GetBiologyId()
	shareBossInfo.deadTime = bossInfo.GetDeadTime()
	shareBossInfo.isDead = bossInfo.GetIsDead()
	return shareBossInfo
}

func convertFromBossInfoList(bossInfoList []*sharebosspb.BossInfo) []*ShareBossInfo {
	shareBossInfoList := make([]*ShareBossInfo, 0, len(bossInfoList))
	for _, bossInfo := range bossInfoList {
		shareBossInfoList = append(shareBossInfoList, convertFromBossInfo(bossInfo))
	}
	return shareBossInfoList
}
