package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droptemplate "fgame/fgame/game/drop/template"
	propertytypes "fgame/fgame/game/property/types"
	xianfuplayer "fgame/fgame/game/xianfu/player"
	xianfutypes "fgame/fgame/game/xianfu/types"
)

func BuildSCXianfuGet(xianfuMap map[xianfutypes.XianfuType]*xianfuplayer.PlayerXianfuObject) *uipb.SCXianfuGet {
	scXianfuGet := &uipb.SCXianfuGet{}
	xfInfoArr := make([]*uipb.XianfuInfo, 0, len(xianfuMap))
	for _, xfObj := range xianfuMap {
		xfInfoArr = append(xfInfoArr, buildXianfuInfo(xfObj))
	}
	scXianfuGet.XianfuInfo = xfInfoArr

	return scXianfuGet
}

func buildXianfuInfo(obj *xianfuplayer.PlayerXianfuObject) *uipb.XianfuInfo {
	xianfuId := obj.GetXianfuId()
	xianfuType := int32(obj.GetXianfuType())
	useTimes := obj.GetUseTimes()
	state := int32(obj.GetState())
	startTime := obj.GetStartTime()
	group := obj.GetGroup()
	xfInfo := &uipb.XianfuInfo{}
	xfInfo.XianfuId = &xianfuId
	xfInfo.XianfuType = &xianfuType
	xfInfo.UseTimes = &useTimes
	xfInfo.State = &state
	xfInfo.StartTime = &startTime
	xfInfo.Group = &group
	return xfInfo
}

func BuildSCXianfuUpgrade(xianfuId int32, xianfuType xianfutypes.XianfuType) *uipb.SCXianfuUpgrade {
	scXianfuUpgrade := &uipb.SCXianfuUpgrade{}
	typ := int32(xianfuType)
	scXianfuUpgrade.XianfuId = &xianfuId
	scXianfuUpgrade.XianfuType = &typ

	return scXianfuUpgrade
}

func BuildSCXianfuAccelerate(xianfuId int32, xianfuType xianfutypes.XianfuType) *uipb.SCXianfuAccelerate {
	scXianfuAccelerate := &uipb.SCXianfuAccelerate{}

	scXianfuAccelerate.XianfuId = &xianfuId
	typ := int32(xianfuType)
	scXianfuAccelerate.XianfuType = &typ

	return scXianfuAccelerate
}

func BuildSCXianfuSaoDang(xianfuId int32, xianfuType xianfutypes.XianfuType, saodangNum int32, totalRewData *propertytypes.RewData, dropItemList [][]*droptemplate.DropItemData) *uipb.SCXianfuSaoDang {
	scXianfuSaoDang := &uipb.SCXianfuSaoDang{}

	scXianfuSaoDang.XianfuId = &xianfuId
	typ := int32(xianfuType)
	scXianfuSaoDang.XianfuType = &typ

	for _, itemList := range dropItemList {
		scXianfuSaoDang.RewAllItemArr = append(scXianfuSaoDang.RewAllItemArr, buildSweepDrop(itemList))
	}
	scXianfuSaoDang.RewProperty = buildRewProperty(totalRewData)
	scXianfuSaoDang.SaodangNum = &saodangNum

	return scXianfuSaoDang
}

func buildRewProperty(rd *propertytypes.RewData) *uipb.RewProperty {
	rewProperty := &uipb.RewProperty{}
	rewExp := rd.GetRewExp()
	rewExpPoint := rd.GetRewExpPoint()
	rewGold := rd.GetRewGold()
	rewBindGold := rd.GetRewBindGold()
	rewSilver := rd.GetRewSilver()

	rewProperty.Exp = &rewExp
	rewProperty.ExpPoint = &rewExpPoint
	rewProperty.Silver = &rewSilver
	rewProperty.Gold = &rewGold
	rewProperty.BindGold = &rewBindGold

	return rewProperty
}

func buildDropInfo(itemId, num, level int32) *uipb.DropInfo {
	dropInfo := &uipb.DropInfo{}
	dropInfo.ItemId = &itemId
	dropInfo.Num = &num
	dropInfo.Level = &level

	return dropInfo
}

func BuildSCXianfuChallenge(xianfuId int32, xianfuType xianfutypes.XianfuType, success bool) *uipb.SCXianfuChallenge {
	scXianfuChallenge := &uipb.SCXianfuChallenge{}

	scXianfuChallenge.XianfuId = &xianfuId
	typ := int32(xianfuType)
	scXianfuChallenge.XianfuType = &typ
	scXianfuChallenge.IsSuccess = &success

	return scXianfuChallenge
}

func BuildSCXianfuChallengeResult(xianfuId int32, xianfuType int32, success bool, resource int64, rewItemMap map[int32]int32, group int32) *uipb.SCXianfuChallengeResult {
	scXianfuChallengeResult := &uipb.SCXianfuChallengeResult{}

	scXianfuChallengeResult.XianfuId = &xianfuId
	scXianfuChallengeResult.XianfuType = &xianfuType
	scXianfuChallengeResult.IsSuccess = &success
	scXianfuChallengeResult.Resource = &resource
	scXianfuChallengeResult.Group = &group
	return scXianfuChallengeResult
}

func BuildSCXianfuSceneInfo(xianfuId, xianfuType, killNum, biologyGroup, totalBiologyGroup int32, res int64, createTime int64) *uipb.SCXianfuSceneInfo {
	scXianfuSceneInfo := &uipb.SCXianfuSceneInfo{}

	scXianfuSceneInfo.XianfuId = &xianfuId
	scXianfuSceneInfo.XianfuType = &xianfuType
	scXianfuSceneInfo.BiologyGroup = &biologyGroup
	scXianfuSceneInfo.TotalBiologyGroup = &totalBiologyGroup
	scXianfuSceneInfo.KillNum = &killNum
	scXianfuSceneInfo.Resource = &res
	scXianfuSceneInfo.CreateTime = &createTime

	return scXianfuSceneInfo
}

func BuildSCXianfuNextBiologyGroup(nextGroup int32) *uipb.SCXianfuRefreshBiology {
	scXianfuRefreshBiology := &uipb.SCXianfuRefreshBiology{}

	scXianfuRefreshBiology.NextBiologyGroup = &nextGroup

	return scXianfuRefreshBiology
}

func BuildSCXianfuFinishAll(xianfuId int32, xianfuType xianfutypes.XianfuType, rd *propertytypes.RewData, itemList [][]*droptemplate.DropItemData) *uipb.SCXianfuFinishAll {
	scXianfuFinishAll := &uipb.SCXianfuFinishAll{}
	scXianfuFinishAll.XianfuId = &xianfuId
	typ := int32(xianfuType)
	scXianfuFinishAll.XianfuType = &typ
	scXianfuFinishAll.RewAllProperty = buildRewProperty(rd)
	for i := int(0); i < len(itemList); i++ {
		scXianfuFinishAll.RewAllItemArr = append(scXianfuFinishAll.RewAllItemArr, buildSweepDrop(itemList[i]))
	}

	return scXianfuFinishAll
}

func buildSweepDrop(dropItemList []*droptemplate.DropItemData) *uipb.XianfuSweepDrop {
	xianfuSweepDrop := &uipb.XianfuSweepDrop{}
	for _, itemData := range dropItemList {
		itemId := itemData.GetItemId()
		num := itemData.GetNum()
		level := itemData.GetLevel()

		xianfuSweepDrop.DropList = append(xianfuSweepDrop.DropList, buildDropInfo(itemId, num, level))
	}
	return xianfuSweepDrop
}

func BuildSCXianfuRewNotice(resource int64) *uipb.SCXianfuRewNotice {
	scXianfuRewNotice := &uipb.SCXianfuRewNotice{}
	scXianfuRewNotice.Num = &resource

	return scXianfuRewNotice
}

func BuildSCXianfuKillNumNotice(num int32) *uipb.SCXianfuKillNumNotice {
	scXianfuKillNumNotice := &uipb.SCXianfuKillNumNotice{}
	scXianfuKillNumNotice.KillNum = &num

	return scXianfuKillNumNotice
}

func BuildSCXianfuBossHpChangedNotice(hp int64) *uipb.SCXianfuBossHpChangedNotice {
	scXianfuBossHpChangedNotice := &uipb.SCXianfuBossHpChangedNotice{}
	scXianfuBossHpChangedNotice.CurHp = &hp

	return scXianfuBossHpChangedNotice
}
