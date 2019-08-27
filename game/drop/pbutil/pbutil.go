package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droptemplate "fgame/fgame/game/drop/template"
)

func BuildSimpleDropInfoList(itemMap map[int32]int32) (dropList []*uipb.DropInfo) {
	for itemId, itemNum := range itemMap {
		dropInfo := BuildSimpleDropInfo(itemId, itemNum)
		dropList = append(dropList, dropInfo)
	}
	return dropList
}

func BuildSimpleDropInfo(itemId, num int32) *uipb.DropInfo {
	dropInfo := &uipb.DropInfo{}
	dropInfo.ItemId = &itemId
	dropInfo.Num = &num
	return dropInfo
}

func BuildDropInfoList(dropDataList []*droptemplate.DropItemData) (dropList []*uipb.DropInfo) {
	for _, itemData := range dropDataList {
		itemId := itemData.ItemId
		num := itemData.Num
		level := itemData.Level
		upstar := itemData.Upstar
		attrList := itemData.AttrList

		dropInfo := BuildDropInfo(itemId, num, level, upstar, attrList)
		dropList = append(dropList, dropInfo)
	}

	return dropList
}

func BuildDropInfo(itemId, num, level, upstar int32, attrList []int32) *uipb.DropInfo {
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
