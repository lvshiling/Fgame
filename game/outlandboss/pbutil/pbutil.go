package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	coretypes "fgame/fgame/core/types"
	commonpbutil "fgame/fgame/game/common/pbutil"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/outlandboss/outlandboss"
	"fgame/fgame/game/scene/scene"
)

func BuildSCOutlandBossList(bossList []scene.NPC, curZhuoQi int32) *uipb.SCOutlandBossList {
	scMsg := &uipb.SCOutlandBossList{}

	for _, boss := range bossList {
		scMsg.BossInfoList = append(scMsg.BossInfoList, buildBossInfo(boss))
	}
	scMsg.CurZhuoQi = &curZhuoQi
	return scMsg
}

func BuildSCOutlandBossChallenge(pos coretypes.Position) *uipb.SCOutlandBossChallenge {
	scMsg := &uipb.SCOutlandBossChallenge{}
	scMsg.Pos = commonpbutil.BuildPos(pos)
	return scMsg
}

func BuildSCOutlandBossInfoBroadcast(boss scene.NPC) *uipb.SCOutlandBossInfoBroadcast {
	scMsg := &uipb.SCOutlandBossInfoBroadcast{}
	scMsg.BossInfo = buildBossInfo(boss)

	return scMsg
}

func BuildSCOutlandBossListInfoNotice(bossList []scene.NPC) *uipb.SCOutlandBossListInfoNotice {
	scMsg := &uipb.SCOutlandBossListInfoNotice{}
	for _, boss := range bossList {
		scMsg.BossInfoList = append(scMsg.BossInfoList, buildBossInfo(boss))
	}
	return scMsg
}

func BuildSCOutlandBossEnemiesNotice(attackId int64, plName, biologyName string, itemList []*droptemplate.DropItemData, biologyId int32) *uipb.SCOutlandBossEnemiesNotice {
	scMsg := &uipb.SCOutlandBossEnemiesNotice{}
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

func BuildSCOutlandBossZhuoqiInfo(zhuoqi int32) *uipb.SCOutlandBossZhuoqiInfo {
	scMsg := &uipb.SCOutlandBossZhuoqiInfo{}
	scMsg.CurZhuoQi = &zhuoqi
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

func buildRecord(obj *outlandboss.OutlandBossDropRecordsObject) *uipb.DropRecords {
	dropRecord := &uipb.DropRecords{}
	plname := obj.GetKillerName()
	biologyId := obj.GetBiologyId()
	mapId := obj.GetMapId()
	dropTime := obj.GetDropTime()

	dropRecord.KillerName = &plname
	dropRecord.BiologyId = &biologyId
	dropRecord.MapId = &mapId
	dropRecord.DropTime = &dropTime

	for itemId, num := range obj.ItemMap {
		dropRecord.ItemList = append(dropRecord.ItemList, buildItem(itemId, num))
	}
	return dropRecord
}

func buildItem(itemId int32, num int32) *uipb.ItemInfo {
	itemInfo := &uipb.ItemInfo{}
	itemInfo.ItemId = &itemId
	itemInfo.Num = &num
	return itemInfo
}

func BuildSCOutlandBossDropRecordsGet(dropDataList []*outlandboss.OutlandBossDropRecordsObject) *uipb.SCOutlandBossDropRecordsGet {
	dropRecords := &uipb.SCOutlandBossDropRecordsGet{}
	for _, item := range dropDataList {
		dropRecords.RecordsList = append(dropRecords.RecordsList, buildRecord(item))
	}
	return dropRecords
}

func BuildSCOutlandBossDropRecordsIncr(dropDataList []*outlandboss.OutlandBossDropRecordsObject) *uipb.SCOutlandBossDropRecordsIncr {
	dropRecords := &uipb.SCOutlandBossDropRecordsIncr{}
	for _, item := range dropDataList {
		dropRecords.RecordsList = append(dropRecords.RecordsList, buildRecord(item))
	}
	return dropRecords
}
