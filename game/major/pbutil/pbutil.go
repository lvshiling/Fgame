package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	droptemplate "fgame/fgame/game/drop/template"
	playermajor "fgame/fgame/game/major/player"
	majortypes "fgame/fgame/game/major/types"
)

func BuildSCMajorScene(starTime int64, ownerId int64, spouseId int64, majorType, fubenId int32) *uipb.SCMajorSceneInfo {
	majorSceneInfo := &uipb.SCMajorSceneInfo{}
	majorSceneInfo.CreateTime = &starTime
	majorSceneInfo.OwnerId = &ownerId
	majorSceneInfo.SpouseId = &spouseId
	majorSceneInfo.MajorType = &majorType
	majorSceneInfo.FubenId = &fubenId
	return majorSceneInfo
}

func BuildSCMajorInvite(inviteTime int64, majorType, fubenId int32) *uipb.SCMajorInvite {
	majorInvite := &uipb.SCMajorInvite{}
	majorInvite.InviteTime = &inviteTime
	majorInvite.MajorType = &majorType
	majorInvite.FubenId = &fubenId
	return majorInvite
}

func BuildSCMajorInvitePushSpouse(playerId int64, majorType, fubenId int32) *uipb.SCMajorInvitePushSpouse {
	majorInvitePushSpouse := &uipb.SCMajorInvitePushSpouse{}
	majorInvitePushSpouse.PlayerId = &playerId
	majorInvitePushSpouse.FubenId = &fubenId
	majorInvitePushSpouse.MajorType = &majorType
	return majorInvitePushSpouse
}

func BuildSCMajorSpouseRefused(name string, majorType, fubenId int32) *uipb.SCMajorSpouseRefused {
	majorSpouseRefused := &uipb.SCMajorSpouseRefused{}
	majorSpouseRefused.Name = &name
	majorSpouseRefused.FubenId = &fubenId
	majorSpouseRefused.MajorType = &majorType
	return majorSpouseRefused
}

func BuildSCMajorInviteDeal(result int32) *uipb.SCMajorInviteDeal {
	majorInviteDeal := &uipb.SCMajorInviteDeal{}
	majorInviteDeal.Result = &result
	return majorInviteDeal
}

func BuildSCMajorInviteCancle(result int32) *uipb.SCMajorInviteCancle {
	majorInviteCancle := &uipb.SCMajorInviteCancle{}
	majorInviteCancle.Result = &result
	return majorInviteCancle
}

func BuildSCMajorInvitePushCancle(name string, majorType, fubenId int32) *uipb.SCMajorInvitePushCancle {
	majorInvitePushCancle := &uipb.SCMajorInvitePushCancle{}
	majorInvitePushCancle.Name = &name
	majorInvitePushCancle.FubenId = &fubenId
	majorInvitePushCancle.MajorType = &majorType
	return majorInvitePushCancle
}

func BuildSCMajorResult(state bool, majorType, fubenId int32) *uipb.SCMajorResult {
	majorResult := &uipb.SCMajorResult{}
	majorResult.Sucess = &state
	majorResult.FubenId = &fubenId
	majorResult.MajorType = &majorType
	return majorResult
}

func BuildSCMajorNum(majorType, num int32) *uipb.SCMajorNum {
	majorNum := &uipb.SCMajorNum{}
	majorNum.Num = &num
	majorNum.MajorType = &majorType
	return majorNum
}

func BuildSCMajorNumNotice(objMap map[majortypes.MajorType]*playermajor.PlayerMajorNumObject) *uipb.SCMajorNumNotice {
	scMsg := &uipb.SCMajorNumNotice{}
	for typ, obj := range objMap {
		majorType := int32(typ)
		num := obj.Num
		scMsg.InfoList = append(scMsg.InfoList, buildMajorNumInfo(majorType, num))

	}
	return scMsg
}

func buildMajorNumInfo(majorType, num int32) *uipb.MajorNumInfo {
	info := &uipb.MajorNumInfo{}
	info.MajorType = &majorType
	info.Num = &num
	return info
}

func BuildSCMajorSaoDang(mType majortypes.MajorType, fubenId int32, saodangNum int32, dropItemList [][]*droptemplate.DropItemData) *uipb.SCMajorSaoDang {
	scMsg := &uipb.SCMajorSaoDang{}
	for _, itemList := range dropItemList {
		scMsg.SweepDropList = append(scMsg.SweepDropList, buildSweepDrop(itemList))
	}
	typ := int32(mType)
	scMsg.MajorType = &typ
	scMsg.FubenId = &fubenId
	scMsg.SaodangNum = &saodangNum

	return scMsg
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
