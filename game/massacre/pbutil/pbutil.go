package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
	uipb "fgame/fgame/common/codec/pb/ui"
	commontypes "fgame/fgame/game/common/types"
	playermassacre "fgame/fgame/game/massacre/player"
	massacretypes "fgame/fgame/game/massacre/types"
)

func BuildSCMassacreGet(bShieldInfo *playermassacre.PlayerMassacreObject) *uipb.SCMassacreGet {
	scMassacreGet := &uipb.SCMassacreGet{}
	scMassacreGet.MassacreInfo = buildObjMassacreInfo(bShieldInfo)
	return scMassacreGet
}

func BuildSCMassacreAdavanced(advancedId int32, sqNum int64, typ commontypes.AdvancedType, resultType int32) *uipb.SCMassacreAdvanced {
	massacreAdvanced := &uipb.SCMassacreAdvanced{}
	massacreAdvanced.AdvancedId = &advancedId
	typInt := int32(typ)
	massacreAdvanced.AdvancedType = &typInt
	massacreAdvanced.ShaQiNum = &sqNum
	massacreAdvanced.ResultType = &resultType

	return massacreAdvanced
}

func BuildSCMassacreInfo(info *massacretypes.MassacreInfo) *uipb.SCMassacreInfo {
	scMassacreInfo := &uipb.SCMassacreInfo{}
	scMassacreInfo.MassacreInfo = buildMassacreInfo(info)
	return scMassacreInfo
}

func BuildSCMassacreWeaponLose(old_id int32, new_id int32, attacker string) *uipb.SCMassacreWeaponLoseInfo {
	massacreWeaponLose := &uipb.SCMassacreWeaponLoseInfo{}
	massacreWeaponLose.OldAdvancedId = &old_id
	massacreWeaponLose.NewAdvancedId = &new_id
	massacreWeaponLose.KillerName = &attacker

	return massacreWeaponLose
}

func BuildSIMassacreDrop(itemId int32, itemNum int64, attackId int64) *crosspb.SIMassacreDrop {
	massacreDrop := &crosspb.SIMassacreDrop{}
	massacreDrop.ItemId = &itemId
	massacreDrop.ItemNum = &itemNum
	massacreDrop.AttackerId = &attackId

	return massacreDrop
}

func BuildSCMassacreShaQiDrop(costStar int32, bagDropNum int32, attackName string) *uipb.SCMassacreShaQiDrop {
	massacreShaQiDrop := &uipb.SCMassacreShaQiDrop{}
	massacreShaQiDrop.DropStar = &costStar
	massacreShaQiDrop.DropNum = &bagDropNum
	massacreShaQiDrop.KillerName = &attackName

	return massacreShaQiDrop
}

func BuildSCMassacreShaQiVary(num int64) *uipb.SCMassacreShaQiVary {
	massacreShaQiVary := &uipb.SCMassacreShaQiVary{}
	massacreShaQiVary.ShaQiNum = &num

	return massacreShaQiVary
}

func buildMassacreInfo(info *massacretypes.MassacreInfo) *uipb.MassacreInfo {
	massacreInfo := &uipb.MassacreInfo{}

	advanceId := info.AdvancedId
	sqNum := info.ShaQiNum
	massacreInfo.ShaQiNum = &sqNum
	massacreInfo.AdvancedId = &advanceId

	return massacreInfo
}

func buildObjMassacreInfo(info *playermassacre.PlayerMassacreObject) *uipb.MassacreInfo {
	massacreInfo := &uipb.MassacreInfo{}
	sqNum := info.ShaQiNum
	advanceId := int32(info.AdvanceId)
	massacreInfo.ShaQiNum = &sqNum
	massacreInfo.AdvancedId = &advanceId

	return massacreInfo
}
