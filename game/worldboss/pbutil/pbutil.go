package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	coretypes "fgame/fgame/core/types"
	commonpbutil "fgame/fgame/game/common/pbutil"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/shareboss/shareboss"
	worldbosstypes "fgame/fgame/game/worldboss/types"
)

func BuildSCWorldBossList(bossList []scene.NPC, bossType int32) *uipb.SCWorldBossList {
	scMsg := buildSCWorldBossList(bossList)
	scMsg.BossType = &bossType
	return scMsg
}

func BuildSCWorldBossListOutBoss(bossList []scene.NPC, zhuoQi, bossType int32) *uipb.SCWorldBossList {
	scMsg := buildSCWorldBossList(bossList)
	scMsg.CurZhuoQi = &zhuoQi
	scMsg.BossType = &bossType
	return scMsg
}

func BuildSCWorldBossListUnrealBoss(bossList []scene.NPC, pilao, buyTimes, bossType int32) *uipb.SCWorldBossList {
	scMsg := buildSCWorldBossList(bossList)
	scMsg.CurPilao = &pilao
	scMsg.CurBuyTimes = &buyTimes
	scMsg.BossType = &bossType
	return scMsg
}

func BuildSCWorldBossListShareBoss(bossList []*shareboss.ShareBossInfo, bossType int32) *uipb.SCWorldBossList {
	scMsg := &uipb.SCWorldBossList{}
	for _, boss := range bossList {
		bossInfo := &uipb.BossInfo{}
		biologyId := boss.GetBiologyId()
		bossInfo.BiologyId = &biologyId
		deadTime := boss.GetDeadTime()
		bossInfo.DeadTime = &deadTime
		isDead := boss.IsDead()
		bossInfo.IsDead = &isDead
		pos := commonpbutil.BuildPos(boss.GetPosition())
		bossInfo.Pos = pos

		scMsg.BossInfoList = append(scMsg.BossInfoList, bossInfo)
	}
	scMsg.BossType = &bossType

	return scMsg
}

func buildSCWorldBossList(bossList []scene.NPC) *uipb.SCWorldBossList {
	scMsg := &uipb.SCWorldBossList{}
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

func BuildSCWorldBossInfoBroadcast(boss scene.NPC, typ worldbosstypes.BossType) *uipb.SCWorldBossInfoBroadcast {
	scWorldBossInfoBroadcast := &uipb.SCWorldBossInfoBroadcast{}
	scWorldBossInfoBroadcast.BossInfo = buildBossInfo(boss)
	typInt := int32(typ)
	scWorldBossInfoBroadcast.BossTyp = &typInt
	return scWorldBossInfoBroadcast
}

func BuildSCWorldBossListInfoNotice(bossList []scene.NPC, typ worldbosstypes.BossType, reliveTime int32) *uipb.SCWorldBossListInfoNotice {
	scWorldBossListInfoNotice := &uipb.SCWorldBossListInfoNotice{}
	for _, boss := range bossList {
		scWorldBossListInfoNotice.BossInfoList = append(scWorldBossListInfoNotice.BossInfoList, buildBossInfo(boss))
	}
	typInt := int32(typ)
	scWorldBossListInfoNotice.BossTyp = &typInt
	scWorldBossListInfoNotice.ReliveTime = &reliveTime
	return scWorldBossListInfoNotice
}

func BuildSCChallengeWorldBoss(pos coretypes.Position, bossType int32) *uipb.SCChallengeWorldBoss {
	scMsg := &uipb.SCChallengeWorldBoss{}
	scMsg.Pos = commonpbutil.BuildPos(pos)
	scMsg.BossType = &bossType
	return scMsg
}

func BuildSCWorldBossReliveTimeNotice(bossType int32, reliveTime int32) *uipb.SCWorldBossReliveTimeNotice {
	scMsg := &uipb.SCWorldBossReliveTimeNotice{}
	scMsg.BossTyp = &bossType
	scMsg.ReliveTime = &reliveTime
	return scMsg
}
