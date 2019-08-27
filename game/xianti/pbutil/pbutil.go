package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	commontypes "fgame/fgame/game/common/types"
	xianticommon "fgame/fgame/game/xianti/common"
	playerxianti "fgame/fgame/game/xianti/player"
	xiantitypes "fgame/fgame/game/xianti/types"
)

func BuildSCXianTiGet(xianTiInfo *playerxianti.PlayerXianTiObject, xianTiOtherMap map[xiantitypes.XianTiType]map[int32]*playerxianti.PlayerXianTiOtherObject) *uipb.SCXiantiGet {
	xianTiGet := &uipb.SCXiantiGet{}
	advancedId := int32(xianTiInfo.AdvanceId)
	xianTiGet.AdvancedId = &advancedId
	xianTiId := xianTiInfo.XianTiId
	xianTiGet.XiantiId = &xianTiId
	unrealLevel := xianTiInfo.UnrealLevel
	unrealPro := xianTiInfo.UnrealPro
	xianTiGet.UnrealLevel = &unrealLevel
	xianTiGet.UnrealPro = &unrealPro
	bless := xianTiInfo.Bless
	xianTiGet.Bless = &bless
	blessTime := xianTiInfo.BlessTime
	xianTiGet.BlessTime = &blessTime
	hidden := false
	if xianTiInfo.Hidden == 1 {
		hidden = true
	}
	xianTiGet.Hidden = &hidden
	for _, unrealId := range xianTiInfo.UnrealList {
		xianTiGet.UnrealList = append(xianTiGet.UnrealList, int32(unrealId))
	}

	for _, xianTiTypeOtherMap := range xianTiOtherMap {
		for _, xianTiOtherObj := range xianTiTypeOtherMap {
			xianTiGet.XiantiSkinList = append(xianTiGet.XiantiSkinList, buildXianTiOther(xianTiOtherObj))
		}
	}
	return xianTiGet
}

func buildXianTiOther(xianTiOtherObj *playerxianti.PlayerXianTiOtherObject) *uipb.XiantiSkinInfo {
	xianTiSkinInfo := &uipb.XiantiSkinInfo{}
	xianTiId := xianTiOtherObj.XianTiId
	level := xianTiOtherObj.Level
	pro := xianTiOtherObj.UpPro

	xianTiSkinInfo.XiantiId = &xianTiId
	xianTiSkinInfo.Level = &level
	xianTiSkinInfo.Pro = &pro
	return xianTiSkinInfo
}

func BuildSCXianTiUnrealDan(level int32, progress int32) *uipb.SCXiantiUnrealDan {
	xianTiUnrealDan := &uipb.SCXiantiUnrealDan{}
	xianTiUnrealDan.Level = &level
	xianTiUnrealDan.Progress = &progress
	return xianTiUnrealDan
}

func BuildSCXianTiUnreal(xianTiId int32) *uipb.SCXiantiUnreal {
	xianTiUnreal := &uipb.SCXiantiUnreal{}
	xianTiUnreal.XiantiId = &xianTiId
	return xianTiUnreal
}

func BuildSCXianTiUnload(xianTiId int32) *uipb.SCXiantiUnload {
	xianTiUnload := &uipb.SCXiantiUnload{}
	xianTiUnload.XiantiId = &xianTiId
	return xianTiUnload
}

func BuildSCXianTiHidden(hiddenFlag bool) *uipb.SCXiantiHidden {
	xianTiHidden := &uipb.SCXiantiHidden{}
	xianTiHidden.Hidden = &hiddenFlag
	return xianTiHidden
}

func BuildSCXianTiAdavancedFinshed(advancedId int32, xianTiId int32, typ commontypes.AdvancedType) *uipb.SCXiantiAdvanced {
	xianTiAdvanced := &uipb.SCXiantiAdvanced{}
	xianTiAdvanced.AdvancedId = &advancedId
	xianTiAdvanced.XiantiId = &xianTiId
	typeInt := int32(typ)
	xianTiAdvanced.AdvancedType = &typeInt
	return xianTiAdvanced
}

func BuildSCXianTiAdavanced(advancedId int32, xianTiId int32, bless int32, totalBless int32, blessTime int64, typ commontypes.AdvancedType, isDouble bool) *uipb.SCXiantiAdvanced {
	xianTiAdvanced := &uipb.SCXiantiAdvanced{}
	xianTiAdvanced.AdvancedId = &advancedId
	xianTiAdvanced.XiantiId = &xianTiId
	xianTiAdvanced.Bless = &bless
	xianTiAdvanced.BlessTime = &blessTime
	typeInt := int32(typ)
	xianTiAdvanced.AdvancedType = &typeInt
	xianTiAdvanced.IsDouble = &isDouble
	xianTiAdvanced.TotalBless = &totalBless
	return xianTiAdvanced
}

func BuildXianTiInfo(info *xianticommon.XianTiInfo) *uipb.XiantiInfo {
	xianTiInfo := &uipb.XiantiInfo{}
	advanceId := int32(info.AdvanceId)
	xianTiInfo.AdvancedId = &advanceId
	xianTiId := info.XianTiId
	xianTiInfo.XiantiId = &xianTiId
	unrealLevel := info.UnrealLevel
	unrealPro := info.UnrealPro
	xianTiInfo.UnrealLevel = &unrealLevel
	xianTiInfo.UnrealPro = &unrealPro
	for _, skinInfo := range info.SkinList {
		temp := buildXianTiSkinInfo(skinInfo)
		xianTiInfo.SkinList = append(xianTiInfo.SkinList, temp)
	}

	return xianTiInfo
}

func buildXianTiSkinInfo(info *xianticommon.XianTiSkinInfo) *uipb.XiantiSkinInfo {
	skinInfo := &uipb.XiantiSkinInfo{}
	xianTiId := info.XianTiId
	level := info.Level
	pro := info.UpPro

	skinInfo.XiantiId = &xianTiId
	skinInfo.Level = &level
	skinInfo.Pro = &pro
	return skinInfo
}

func BuildSCXianTiUpstar(xianTiId int32, level int32, pro int32) *uipb.SCXiantiUpstar {
	xianTiUpstar := &uipb.SCXiantiUpstar{}
	xianTiUpstar.XiantiId = &xianTiId
	xianTiUpstar.Level = &level
	xianTiUpstar.UpPro = &pro
	return xianTiUpstar
}
