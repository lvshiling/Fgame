package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	commontypes "fgame/fgame/game/common/types"
	fabaocommon "fgame/fgame/game/fabao/common"
	playerfabao "fgame/fgame/game/fabao/player"
	fabaotypes "fgame/fgame/game/fabao/types"
)

func BuildSCFaBaoGet(faBaoInfo *playerfabao.PlayerFaBaoObject, fabaoOtherMap map[fabaotypes.FaBaoType]map[int32]*playerfabao.PlayerFaBaoOtherObject) *uipb.SCFaBaoGet {
	faBaoGet := &uipb.SCFaBaoGet{}
	faBaoGet.FaBaoInfo = buildFaBao(faBaoInfo)

	for _, faBaoTypeOtherMap := range fabaoOtherMap {
		for _, faBaoOtherObj := range faBaoTypeOtherMap {
			faBaoGet.FaBaoSkinList = append(faBaoGet.FaBaoSkinList, buildFaBaoOther(faBaoOtherObj))
		}
	}
	return faBaoGet
}

func buildFaBao(faBaoObj *playerfabao.PlayerFaBaoObject) *uipb.FaBaoInfo {
	faBaoInfo := &uipb.FaBaoInfo{}
	advancedId := faBaoObj.GetAdvancedId()
	faBaoId := faBaoObj.GetFaBaoId()
	unrealLevel := faBaoObj.GetUnrealLevel()
	unrealPro := faBaoObj.GetUnrealPro()
	bless := faBaoObj.GetBless()
	blessTime := faBaoObj.GetBlessTime()
	tonglingLevel := faBaoObj.GetTongLingLevel()
	tonglingNum := faBaoObj.GetTongLingNum()
	tonglingPro := faBaoObj.GetTongLingPro()
	hidden := false
	if faBaoObj.GetHidden() == 1 {
		hidden = true
	}

	for _, unrealId := range faBaoObj.GetUnrealList() {
		faBaoInfo.UnrealList = append(faBaoInfo.UnrealList, int32(unrealId))
	}

	faBaoInfo.AdvancedId = &advancedId
	faBaoInfo.FaBaoId = &faBaoId
	faBaoInfo.UnrealLevel = &unrealLevel
	faBaoInfo.UnrealPro = &unrealPro
	faBaoInfo.Bless = &bless
	faBaoInfo.BlessTime = &blessTime
	faBaoInfo.TonglingLevel = &tonglingLevel
	faBaoInfo.TonglingNum = &tonglingNum
	faBaoInfo.TonglingPro = &tonglingPro
	faBaoInfo.Hidden = &hidden
	return faBaoInfo
}

func buildFaBaoOther(faBaoOtherObj *playerfabao.PlayerFaBaoOtherObject) *uipb.FaBaoSkinInfo {
	faBaoSkinInfo := &uipb.FaBaoSkinInfo{}
	faBaoId := faBaoOtherObj.GetFaBaoId()
	level := faBaoOtherObj.GetLevel()
	pro := faBaoOtherObj.GetUpPro()

	faBaoSkinInfo.FaBaoId = &faBaoId
	faBaoSkinInfo.Level = &level
	faBaoSkinInfo.Pro = &pro
	return faBaoSkinInfo
}

func BuildSCFaBaoUnrealDan(level int32, progress int32) *uipb.SCFaBaoUnrealDan {
	faBaoUnrealDan := &uipb.SCFaBaoUnrealDan{}
	faBaoUnrealDan.Level = &level
	faBaoUnrealDan.Progress = &progress
	return faBaoUnrealDan
}

func BuildSCFaBaoUnreal(faBaoId int32) *uipb.SCFaBaoUnreal {
	faBaoUnreal := &uipb.SCFaBaoUnreal{}
	faBaoUnreal.FaBaoId = &faBaoId
	return faBaoUnreal
}

func BuildSCFaBaoUnload(faBaoId int32) *uipb.SCFaBaoUnload {
	unLoad := &uipb.SCFaBaoUnload{}
	unLoad.FaBaoId = &faBaoId
	return unLoad
}

func BuildSCFaBaoAdavancedFinshed(advancedId int32, faBaoId int32, typ commontypes.AdvancedType) *uipb.SCFaBaoAdvanced {
	faBaoAdvanced := &uipb.SCFaBaoAdvanced{}
	faBaoAdvanced.AdvancedId = &advancedId
	faBaoAdvanced.FaBaoId = &faBaoId
	typInt := int32(typ)
	faBaoAdvanced.AdvancedType = &typInt
	return faBaoAdvanced
}

func BuildSCFaBaoAdavanced(advancedId int32, faBaoId int32, bless int32, totalBless int32, blessTime int64, typ commontypes.AdvancedType, isDouble bool) *uipb.SCFaBaoAdvanced {
	faBaoAdvanced := &uipb.SCFaBaoAdvanced{}
	faBaoAdvanced.AdvancedId = &advancedId
	faBaoAdvanced.FaBaoId = &faBaoId
	faBaoAdvanced.Bless = &bless
	faBaoAdvanced.BlessTime = &blessTime
	typInt := int32(typ)
	faBaoAdvanced.AdvancedType = &typInt
	faBaoAdvanced.IsDouble = &isDouble
	faBaoAdvanced.TotalBless = &totalBless
	return faBaoAdvanced
}

func BuildFaBaoInfo(info *fabaocommon.FaBaoInfo) *uipb.FaBaoCacheInfo {
	faBaoCacheInfo := &uipb.FaBaoCacheInfo{}
	advancedId := info.AdvanceId
	unrealLevel := info.UnrealLevel
	unrealPro := info.UnrealPro
	faBaoId := info.FaBaoId
	tonglingLevel := info.TongLingLevel
	tonglingPro := info.TongLingPro

	faBaoCacheInfo.FaBaoId = &faBaoId
	faBaoCacheInfo.AdvancedId = &advancedId
	faBaoCacheInfo.UnrealLevel = &unrealLevel
	faBaoCacheInfo.UnrealPro = &unrealPro
	faBaoCacheInfo.TonglingLevel = &tonglingLevel
	faBaoCacheInfo.TonglingPro = &tonglingPro
	for _, skinInfo := range info.SkinList {
		temp := buildFaBaoSkinInfo(skinInfo)
		faBaoCacheInfo.SkinList = append(faBaoCacheInfo.SkinList, temp)
	}
	return faBaoCacheInfo
}

func buildFaBaoSkinInfo(info *fabaocommon.FaBaoSkinInfo) *uipb.FaBaoSkinInfo {
	skinInfo := &uipb.FaBaoSkinInfo{}
	faBaoId := info.FaBaoId
	level := info.Level
	pro := info.UpPro

	skinInfo.FaBaoId = &faBaoId
	skinInfo.Level = &level
	skinInfo.Pro = &pro
	return skinInfo
}

func BuildSCFaBaoHidden(hiddenFlag bool) *uipb.SCFaBaoHidden {
	faBaoHidden := &uipb.SCFaBaoHidden{}
	faBaoHidden.Hidden = &hiddenFlag
	return faBaoHidden
}

func BuildSCFaBaoTongLing(level int32, num int32, pro int32) *uipb.SCFaBaoTongLing {
	scFaBaoTongLing := &uipb.SCFaBaoTongLing{}
	scFaBaoTongLing.TonglingLevel = &level
	scFaBaoTongLing.TonglingNum = &num
	scFaBaoTongLing.TonglingPro = &pro
	return scFaBaoTongLing
}

func BuildSCFaBaoUpstar(faBaoId int32, level int32, pro int32) *uipb.SCFaBaoUpstar {
	faBaoUpstar := &uipb.SCFaBaoUpstar{}
	faBaoUpstar.FaBaoId = &faBaoId
	faBaoUpstar.Level = &level
	faBaoUpstar.UpPro = &pro
	return faBaoUpstar
}
