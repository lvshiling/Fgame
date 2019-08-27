package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	coretypes "fgame/fgame/core/types"
	commonpbutil "fgame/fgame/game/common/pbutil"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/scene/scene"
)

func BuildSCUnrealBossList(bossList []scene.NPC, curPilao, curBuyTimes int32) *uipb.SCUnrealBossList {
	scMsg := &uipb.SCUnrealBossList{}

	for _, boss := range bossList {
		scMsg.BossInfoList = append(scMsg.BossInfoList, buildBossInfo(boss))
	}
	scMsg.CurPilao = &curPilao
	scMsg.CurBuyTimes = &curBuyTimes
	return scMsg
}

func BuildSCUnrealBossChallenge(pos coretypes.Position) *uipb.SCUnrealBossChallenge {
	scMsg := &uipb.SCUnrealBossChallenge{}
	scMsg.Pos = commonpbutil.BuildPos(pos)
	return scMsg
}

func BuildSCUnrealBossInfoBroadcast(boss scene.NPC) *uipb.SCUnrealBossInfoBroadcast {
	scMsg := &uipb.SCUnrealBossInfoBroadcast{}
	scMsg.BossInfo = buildBossInfo(boss)

	return scMsg
}

func BuildSCUnrealBossListInfoNotice(bossList []scene.NPC) *uipb.SCUnrealBossListInfoNotice {
	scMsg := &uipb.SCUnrealBossListInfoNotice{}
	for _, boss := range bossList {
		scMsg.BossInfoList = append(scMsg.BossInfoList, buildBossInfo(boss))
	}
	return scMsg
}

func BuildSCUnrealBossBuyPilaoNum(pilao int32) *uipb.SCUnrealBossBuyPilaoNum {
	scMsg := &uipb.SCUnrealBossBuyPilaoNum{}
	scMsg.CurPilao = &pilao
	return scMsg
}

func BuildSCUnrealBossEnemiesNotice(attackId int64, plName, biologyName string, itemList []*droptemplate.DropItemData, biologyId int32) *uipb.SCUnrealBossEnemiesNotice {
	scMsg := &uipb.SCUnrealBossEnemiesNotice{}
	scMsg.AttackPlayerName = &plName
	scMsg.BossName = &biologyName
	scMsg.AttackId = &attackId
	scMsg.BiologyId = &biologyId
	for _, itemData := range itemList {
		itemId := itemData.ItemId
		num := itemData.Num
		level := itemData.Level
		upstar := itemData.Upstar
		attrList := itemData.AttrList

		scMsg.DropInfoList = append(scMsg.DropInfoList, buildDropInfo(itemId, num, level, upstar, attrList))
	}

	return scMsg
}

func BuildSCUnrealBossPilaoInfo(pilao, curBuyTimes int32) *uipb.SCUnrealBossPilaoInfo {
	scMsg := &uipb.SCUnrealBossPilaoInfo{}
	scMsg.CurPilao = &pilao
	scMsg.CurBuyTimes = &curBuyTimes
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

func buildDropInfo(itemId, num, level, upstar int32, attrList []int32) *uipb.DropInfo {
	dropInfo := &uipb.DropInfo{}
	dropInfo.ItemId = &itemId
	dropInfo.Num = &num
	dropInfo.Level = &level
	dropInfo.UpstarLevel = &upstar
	dropInfo.Attr = buildGoldEquipAttrList(attrList)

	return dropInfo
}

func buildGoldEquipAttrList(attrList []int32) *uipb.GoldEquipAttrList {
	info := &uipb.GoldEquipAttrList{}
	info.AttrList = attrList
	return info
}
