package grpc_pbutil

import (
	coretypes "fgame/fgame/core/types"
	sharebosspb "fgame/fgame/cross/shareboss/pb"
	"fgame/fgame/game/scene/scene"
)

func BuildPos(pos coretypes.Position) *sharebosspb.Position {
	tpos := &sharebosspb.Position{}
	tpos.X = float32(pos.X)
	tpos.Y = float32(pos.Y)
	tpos.Z = float32(pos.Z)
	return tpos
}

func BuildBossInfo(boss scene.NPC) *sharebosspb.BossInfo {
	bossInfo := &sharebosspb.BossInfo{}
	bossInfo.BiologyId = int32(boss.GetBiologyTemplate().TemplateId())
	bossInfo.IsDead = boss.IsDead()
	bossInfo.DeadTime = boss.GetDeadTime()
	bossInfo.Pos = BuildPos(boss.GetPosition())
	return bossInfo
}

func BuildBossInfoList(bossList []scene.NPC) []*sharebosspb.BossInfo {
	bossInfoList := make([]*sharebosspb.BossInfo, 0, len(bossList))
	for _, boss := range bossList {
		bossInfo := BuildBossInfo(boss)
		bossInfoList = append(bossInfoList, bossInfo)
	}
	return bossInfoList
}
