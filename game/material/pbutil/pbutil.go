package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droptemplate "fgame/fgame/game/drop/template"
	materialplayer "fgame/fgame/game/material/player"
	materialtypes "fgame/fgame/game/material/types"
)

func BuildSCMaterialInfoGet(materialMap map[materialtypes.MaterialType]*materialplayer.PlayerMaterialObject) *uipb.SCMaterialInfoGet {
	scMsg := &uipb.SCMaterialInfoGet{}
	for _, obj := range materialMap {
		scMsg.InfoList = append(scMsg.InfoList, buildMaterialInfo(obj))
	}

	return scMsg
}

func BuildSCMaterialChallenge(typ materialtypes.MaterialType, success bool) *uipb.SCMaterialChallenge {
	scMsg := &uipb.SCMaterialChallenge{}
	typInt := int32(typ)
	scMsg.MaterialType = &typInt
	scMsg.IsSuccess = &success

	return scMsg
}

func BuildSCMaterialSceneInfo(materialType, biologyGroup, totalBiologyGroup int32, createTime int64) *uipb.SCMaterialSceneInfo {
	scMsg := &uipb.SCMaterialSceneInfo{}

	scMsg.MaterialType = &materialType
	scMsg.BiologyGroup = &biologyGroup
	scMsg.TotalBiologyGroup = &totalBiologyGroup
	scMsg.CreateTime = &createTime

	return scMsg
}

func BuildSCMaterialRefreshBiology(nextGroup int32) *uipb.SCMaterialRefreshBiology {
	scMsg := &uipb.SCMaterialRefreshBiology{}
	scMsg.NextBiologyGroup = &nextGroup
	return scMsg
}

func BuildSCMaterialSaoDang(materialType materialtypes.MaterialType, saodangNum int32, dropItemList [][]*droptemplate.DropItemData) *uipb.SCMaterialSaoDang {
	scMsg := &uipb.SCMaterialSaoDang{}

	for _, itemList := range dropItemList {
		scMsg.SweepDropList = append(scMsg.SweepDropList, buildSweepDrop(itemList))
	}
	typ := int32(materialType)
	scMsg.MaterialType = &typ
	scMsg.SaodangNum = &saodangNum

	return scMsg
}

func BuildSCMaterialChallengeResult(materialType int32, success bool, rewItemMap map[int32]int32, group int32) *uipb.SCMaterialChallengeResult {
	scMsg := &uipb.SCMaterialChallengeResult{}

	scMsg.MaterialType = &materialType
	scMsg.IsSuccess = &success
	scMsg.Group = &group
	for itemId, num := range rewItemMap {
		level := int32(0)
		scMsg.DropList = append(scMsg.DropList, buildDropInfo(itemId, num, level))
	}
	return scMsg
}

func buildMaterialInfo(obj *materialplayer.PlayerMaterialObject) *uipb.MaterialInfo {
	typInt := int32(obj.GetMaterialType())
	useTimes := obj.GetUseTimes()
	group := obj.GetGroup()

	info := &uipb.MaterialInfo{}
	info.MaterialType = &typInt
	info.UseTimes = &useTimes
	info.Group = &group
	return info
}

func buildSweepDrop(dropItemList []*droptemplate.DropItemData) *uipb.MaterialSweepDrop {
	materialSweepDrop := &uipb.MaterialSweepDrop{}
	for _, itemData := range dropItemList {
		itemId := itemData.GetItemId()
		num := itemData.GetNum()
		level := itemData.GetLevel()
		materialSweepDrop.DropList = append(materialSweepDrop.DropList, buildDropInfo(itemId, num, level))
	}
	return materialSweepDrop
}

func buildDropInfo(itemId, num, level int32) *uipb.DropInfo {
	dropInfo := &uipb.DropInfo{}
	dropInfo.ItemId = &itemId
	dropInfo.Num = &num
	dropInfo.Level = &level

	return dropInfo
}

// func BuildSCMaterialFinishAll(materialId int32, materialType materialtypes.MaterialType, rd *propertytypes.RewData, itemList [][]*droptemplate.DropItemData) *uipb.SCMaterialFinishAll {
// 	scMaterialFinishAll := &uipb.SCMaterialFinishAll{}
// 	scMaterialFinishAll.MaterialId = &materialId
// 	typ := int32(materialType)
// 	scMaterialFinishAll.MaterialType = &typ
// 	scMaterialFinishAll.RewAllProperty = buildRewProperty(rd)
// 	for i := int(0); i < len(itemList); i++ {
// 		scMaterialFinishAll.RewAllItemArr = append(scMaterialFinishAll.RewAllItemArr, buildSweepDrop(itemList[i]))
// 	}

// 	return scMaterialFinishAll
// }
