package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	coretypes "fgame/fgame/core/types"
	commonpbutil "fgame/fgame/game/common/pbutil"
	"fgame/fgame/game/scene/scene"
)

func BuildSCCangJingGeBossList(bossList []scene.NPC) *uipb.SCCangjinggeBossList {
	scMsg := &uipb.SCCangjinggeBossList{}

	for _, boss := range bossList {
		scMsg.BossInfoList = append(scMsg.BossInfoList, buildBossInfo(boss))
	}
	return scMsg
}

func BuildSCCangJingGeBossChallenge(pos coretypes.Position) *uipb.SCCangjinggeBossChallenge {
	scMsg := &uipb.SCCangjinggeBossChallenge{}
	scMsg.Pos = commonpbutil.BuildPos(pos)
	return scMsg
}

func BuildSCCangJingGeBossInfoBroadcast(boss scene.NPC) *uipb.SCCangjinggeBossInfoBroadcast {
	scMsg := &uipb.SCCangjinggeBossInfoBroadcast{}
	scMsg.BossInfo = buildBossInfo(boss)

	return scMsg
}

func BuildSCCangJingGeBossListInfoNotice(bossList []scene.NPC) *uipb.SCCangjinggeBossListInfoNotice {
	scMsg := &uipb.SCCangjinggeBossListInfoNotice{}
	for _, boss := range bossList {
		scMsg.BossInfoList = append(scMsg.BossInfoList, buildBossInfo(boss))
	}
	return scMsg
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
