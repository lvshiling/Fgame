package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droptemplate "fgame/fgame/game/drop/template"
)

func BuildSCMyBossChallenge() *uipb.SCMyBossChallenge {
	scMsg := &uipb.SCMyBossChallenge{}
	return scMsg
}

func BuildSCMyBossInfoNotice(timesMap map[int32]int32) *uipb.SCMyBossInfoNotice {
	scMsg := &uipb.SCMyBossInfoNotice{}
	for biologyId, times := range timesMap {
		scMsg.InfoList = append(scMsg.InfoList, buildMyBossInfo(biologyId, times))
	}
	return scMsg
}

func BuildSCMyBossChallengeResult(isSuccess bool, itemList []*droptemplate.DropItemData) *uipb.SCMyBossChallengeResult {
	scMsg := &uipb.SCMyBossChallengeResult{}
	scMsg.IsSuccess = &isSuccess

	for _, itemData := range itemList {
		itemId := itemData.ItemId
		num := itemData.Num
		level := itemData.Level
		upstar := itemData.Upstar
		attrList := itemData.AttrList

		scMsg.DropList = append(scMsg.DropList, buildDropInfo(itemId, num, level, upstar, attrList))
	}
	return scMsg
}

func BuildSCMyBossSceneInfo(createTime int64, bossId int32) *uipb.SCMyBossSceneInfo {
	scMsg := &uipb.SCMyBossSceneInfo{}
	scMsg.CreateTime = &createTime
	scMsg.BossId = &bossId
	return scMsg
}

func buildMyBossInfo(biologyId int32, times int32) *uipb.MyBossInfo {
	info := &uipb.MyBossInfo{}
	info.BiologyId = &biologyId
	info.AttendTimes = &times
	return info
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
