package pbuitl

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droptemplate "fgame/fgame/game/drop/template"
	"fgame/fgame/game/emperor/emperor"
)

func BuildSCEmperorGet(emperorObj *emperor.EmperorObject, worshipNum int32) *uipb.SCEmperorGet {
	emperorGet := &uipb.SCEmperorGet{}
	emperorGet.EmperorInfo = buildEmperor(emperorObj)
	emperorGet.WorshipNum = &worshipNum
	return emperorGet
}

func BuildSCEmperorWorship(num int32, storage int64) *uipb.SCEmperorWorship {
	emperorWorship := &uipb.SCEmperorWorship{}
	emperorWorship.Num = &num
	emperorWorship.Storage = &storage
	return emperorWorship
}

func BuildSCEmperorStorageGet(emperorObj *emperor.EmperorObject, success bool) *uipb.SCEmperorStorageGet {
	emperorStorageGet := &uipb.SCEmperorStorageGet{}
	emperorStorageGet.Success = &success
	emperorStorageGet.EmperorInfo = buildEmperor(emperorObj)
	return emperorStorageGet
}

func BuildSCEmperorRob(emperorObj *emperor.EmperorObject, dropItemList []*droptemplate.DropItemData, success bool) *uipb.SCEmperorRob {
	emperorRob := &uipb.SCEmperorRob{}
	emperorRob.Success = &success
	for _, dropItem := range dropItemList {
		itemId := dropItem.GetItemId()
		num := dropItem.GetNum()
		emperorRob.ItemList = append(emperorRob.ItemList, buildItem(itemId, num))
	}
	emperorRob.EmperorInfo = buildEmperor(emperorObj)
	return emperorRob
}

func BuildSCEmperorRecordsGet(robList []*emperor.EmperorRecordsObject) *uipb.SCEmperorRecords {
	emperorRecords := &uipb.SCEmperorRecords{}
	for _, rob := range robList {
		emperorRecords.RecordList = append(emperorRecords.RecordList, buildRecord(rob))
	}
	return emperorRecords
}

func BuildSCEmperorRobbed() *uipb.SCEmperorRobbed {
	emperorRobbed := &uipb.SCEmperorRobbed{}
	return emperorRobbed
}

func BuildSCEmperorOpenBox(emperorObj *emperor.EmperorObject, dropItemList []*droptemplate.DropItemData) *uipb.SCEmperorOPenBox {
	emperorOPenBox := &uipb.SCEmperorOPenBox{}
	emperorOPenBox.EmperorInfo = buildEmperor(emperorObj)
	for _, dropItem := range dropItemList {
		itemId := dropItem.GetItemId()
		num := dropItem.GetNum()
		emperorOPenBox.ItemList = append(emperorOPenBox.ItemList, buildItem(itemId, num))
	}
	return emperorOPenBox
}

func buildEmperor(emperorObj *emperor.EmperorObject) *uipb.EmperorInfo {
	emperorInfo := &uipb.EmperorInfo{}
	emperorId := emperorObj.EmperorId
	name := emperorObj.Name
	sex := int32(emperorObj.Sex)
	spouseName := emperorObj.SpouseName
	robNum := emperorObj.RobNum
	robTime := emperorObj.RobTime
	storage := emperorObj.Storage
	lastTime := emperorObj.LastTime
	boxNum := emperorObj.BoxNum
	boxLastTime := emperorObj.BoxLastTime

	emperorInfo.EmperorId = &emperorId
	emperorInfo.Name = &name
	emperorInfo.Sex = &sex
	emperorInfo.SpouseName = &spouseName
	emperorInfo.Storage = &storage
	emperorInfo.RobNum = &robNum
	emperorInfo.RobTime = &robTime
	emperorInfo.LastTime = &lastTime
	emperorInfo.BoxNum = &boxNum
	emperorInfo.BoxLastTime = &boxLastTime
	return emperorInfo
}

func buildRecord(rob *emperor.EmperorRecordsObject) *uipb.EmperorRecord {
	emperorRecord := &uipb.EmperorRecord{}
	typ := rob.Type
	name := rob.EmperorName
	robbedName := rob.RobbedName
	robTime := rob.RobTime
	emperorRecord.Type = &typ
	emperorRecord.Name = &name
	emperorRecord.RobbedName = &robbedName
	emperorRecord.RobTime = &robTime

	for itemId, num := range rob.ItemMap {
		emperorRecord.ItemList = append(emperorRecord.ItemList, buildItem(itemId, num))
	}
	return emperorRecord
}

func buildItem(itemId int32, num int32) *uipb.ItemInfo {
	itemInfo := &uipb.ItemInfo{}
	itemInfo.ItemId = &itemId
	itemInfo.Num = &num
	return itemInfo
}
