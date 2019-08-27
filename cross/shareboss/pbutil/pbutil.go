package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	coretypes "fgame/fgame/core/types"
	commonpbutil "fgame/fgame/game/common/pbutil"
	"fgame/fgame/game/scene/scene"
)

func BuildSCShareBossList(bossList []scene.NPC) *uipb.SCShareBossList {
	scShareBossList := &uipb.SCShareBossList{}

	for _, boss := range bossList {
		scShareBossList.BossInfoList = append(scShareBossList.BossInfoList, buildBossInfo(boss))
	}

	return scShareBossList
}

func buildBossInfo(boss scene.NPC) *uipb.BossInfo {
	bossInfo := &uipb.BossInfo{}
	biologyId := int32(boss.GetBiologyTemplate().TemplateId())
	bossInfo.BiologyId = &biologyId
	deadTime := boss.GetDeadTime()
	bossInfo.DeadTime = &deadTime
	isDead := boss.IsDead()
	bossInfo.IsDead = &isDead
	pos := commonpbutil.BuildPos(boss.GetPosition())
	bossInfo.Pos = pos

	return bossInfo
}

func BuildSCShareBossInfoBroadcast(boss scene.NPC) *uipb.SCShareBossInfoBroadcast {
	scShareBossInfoBroadcast := &uipb.SCShareBossInfoBroadcast{}
	scShareBossInfoBroadcast.BossInfo = buildBossInfo(boss)
	return scShareBossInfoBroadcast
}

func BuildSCShareBossListInfoNotice(bossList []scene.NPC) *uipb.SCShareBossListInfoNotice {
	scShareBossListInfoNotice := &uipb.SCShareBossListInfoNotice{}
	for _, boss := range bossList {
		scShareBossListInfoNotice.BossInfoList = append(scShareBossListInfoNotice.BossInfoList, buildBossInfo(boss))
	}
	return scShareBossListInfoNotice
}

func BuildSCShareBossChallenge(pos coretypes.Position) *uipb.SCShareBossChallenge {
	scShareBossChallenge := &uipb.SCShareBossChallenge{}
	scShareBossChallenge.Pos = commonpbutil.BuildPos(pos)
	return scShareBossChallenge
}
