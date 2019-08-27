package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/treasurebox/treasurebox"
)

func BuildSCInventoryBoxDropInfo(dropItemList []*droptemplate.DropItemData) *uipb.SCInventoryBoxDropInfo {
	scInventoryBoxDropInfo := &uipb.SCInventoryBoxDropInfo{}
	for _, itemData := range dropItemList {
		itemId := itemData.GetItemId()
		num := itemData.GetNum()
		level := itemData.GetLevel()
		upstar := itemData.Upstar
		attrList := itemData.AttrList

		scInventoryBoxDropInfo.DropInfoList = append(scInventoryBoxDropInfo.DropInfoList, buildDropInfo(itemId, num, level, upstar, attrList))
	}

	return scInventoryBoxDropInfo
}

func BuildSCTreasureBoxLog(boxLogInfoList []*treasurebox.BoxLogInfo) *uipb.SCTreasureBoxLog {
	treasureBoxLog := &uipb.SCTreasureBoxLog{}

	for _, boxLogInfo := range boxLogInfoList {
		treasureBoxLog.BoxLogList = append(treasureBoxLog.BoxLogList, buildBoxLog(boxLogInfo))
	}
	return treasureBoxLog
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

func buildBoxLog(boxLogInfo *treasurebox.BoxLogInfo) *uipb.TreasureBoxLog {
	treasureBoxLog := &uipb.TreasureBoxLog{}
	serverId := boxLogInfo.ServerId
	playerName := boxLogInfo.PlayerName
	lastTime := boxLogInfo.LastTime
	treasureBoxLog.ServerId = &serverId
	treasureBoxLog.PlayerName = &playerName
	treasureBoxLog.LastTime = &lastTime

	for itemId, num := range boxLogInfo.ItemMap {
		treasureBoxLog.ItemList = append(treasureBoxLog.ItemList, buildItem(itemId, num))
	}

	return treasureBoxLog
}

func buildItem(itemId int32, num int32) *uipb.ItemInfo {
	itmeInfo := &uipb.ItemInfo{}
	itmeInfo.ItemId = &itemId
	itmeInfo.Num = &num
	return itmeInfo
}

func buildGoldEquipAttrList(attrList []int32) *uipb.GoldEquipAttrList {
	info := &uipb.GoldEquipAttrList{}
	info.AttrList = attrList
	return info
}
