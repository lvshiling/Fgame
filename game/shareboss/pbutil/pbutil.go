package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	coretypes "fgame/fgame/core/types"
	commonpbutil "fgame/fgame/game/common/pbutil"
	"fgame/fgame/game/shareboss/shareboss"
)

func BuildSCShareBossList(bossList []*shareboss.ShareBossInfo) *uipb.SCShareBossList {
	scShareBossList := &uipb.SCShareBossList{}

	for _, boss := range bossList {
		scShareBossList.BossInfoList = append(scShareBossList.BossInfoList, buildBossInfo(boss))
	}

	return scShareBossList
}

func buildBossInfo(boss *shareboss.ShareBossInfo) *uipb.BossInfo {
	bossInfo := &uipb.BossInfo{}
	biologyId := boss.GetBiologyId()
	bossInfo.BiologyId = &biologyId
	deadTime := boss.GetDeadTime()
	bossInfo.DeadTime = &deadTime
	isDead := boss.IsDead()
	bossInfo.IsDead = &isDead
	pos := commonpbutil.BuildPos(boss.GetPosition())
	bossInfo.Pos = pos

	return bossInfo
}

func BuildSCShareBossChallenge(pos coretypes.Position) *uipb.SCShareBossChallenge {
	scShareBossChallenge := &uipb.SCShareBossChallenge{}
	scShareBossChallenge.Pos = commonpbutil.BuildPos(pos)
	return scShareBossChallenge
}
