package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	commontypes "fgame/fgame/game/common/types"
	playershenfa "fgame/fgame/game/shenfa/player"
	shenfatypes "fgame/fgame/game/shenfa/types"
)

func BuildSCShenfaGet(shenfaInfo *playershenfa.PlayerShenfaObject, shenFaOtherMap map[shenfatypes.ShenfaType]map[int32]*playershenfa.PlayerShenfaOtherObject) *uipb.SCShenfaGet {
	shenfaGet := &uipb.SCShenfaGet{}
	advancedId := int32(shenfaInfo.AdvanceId)
	shenfaGet.AdvancedId = &advancedId
	shenfaId := shenfaInfo.ShenfaId
	shenfaGet.ShenfaId = &shenfaId
	unrealLevel := shenfaInfo.UnrealLevel
	unrealPro := shenfaInfo.UnrealPro
	shenfaGet.UnrealLevel = &unrealLevel
	shenfaGet.UnrealPro = &unrealPro
	bless := shenfaInfo.Bless
	shenfaGet.Bless = &bless
	blessTime := shenfaInfo.BlessTime
	shenfaGet.BlessTime = &blessTime
	hidden := false
	if shenfaInfo.Hidden == 1 {
		hidden = true
	}
	shenfaGet.Hidden = &hidden
	for _, unrealId := range shenfaInfo.UnrealList {
		shenfaGet.UnrealList = append(shenfaGet.UnrealList, int32(unrealId))
	}

	for _, shenFaTypeOtherMap := range shenFaOtherMap {
		for _, shenFaOtherObj := range shenFaTypeOtherMap {
			shenfaGet.ShenFaSkinList = append(shenfaGet.ShenFaSkinList, buildShenFaOther(shenFaOtherObj))
		}
	}
	return shenfaGet
}

func buildShenFaOther(shenFaOtherObj *playershenfa.PlayerShenfaOtherObject) *uipb.ShenFaSkinInfo {
	shenFaSkinInfo := &uipb.ShenFaSkinInfo{}
	shenFaId := shenFaOtherObj.ShenFaId
	level := shenFaOtherObj.Level
	pro := shenFaOtherObj.UpPro

	shenFaSkinInfo.ShenFaId = &shenFaId
	shenFaSkinInfo.Level = &level
	shenFaSkinInfo.Pro = &pro
	return shenFaSkinInfo
}

func BuildSCShenfaUnrealDan(level int32, progress int32) *uipb.SCShenfaUnrealDan {
	shenfaUnrealDan := &uipb.SCShenfaUnrealDan{}
	shenfaUnrealDan.Level = &level
	shenfaUnrealDan.Progress = &progress
	return shenfaUnrealDan
}

func BuildSCShenfaUnreal(shenfaId int32) *uipb.SCShenfaUnreal {
	shenfaUnreal := &uipb.SCShenfaUnreal{}
	shenfaUnreal.ShenfaId = &shenfaId
	return shenfaUnreal
}

func BuildSCShenfaUnload(shenfaId int32) *uipb.SCShenfaUnload {
	unLoad := &uipb.SCShenfaUnload{}
	unLoad.ShenfaId = &shenfaId
	return unLoad
}

func BuildSCShenfaAdavancedFinshed(advancedId int32, shenfaId int32, typ commontypes.AdvancedType) *uipb.SCShenfaAdvanced {
	shenfaAdvanced := &uipb.SCShenfaAdvanced{}
	shenfaAdvanced.AdvancedId = &advancedId
	shenfaAdvanced.ShenfaId = &shenfaId
	typeInt := int32(typ)
	shenfaAdvanced.AdvancedType = &typeInt
	return shenfaAdvanced
}

func BuildSCShenfaAdavanced(advancedId int32, shenfaId int32, bless int32, totalBless int32, blessTime int64, typ commontypes.AdvancedType, isDouble bool) *uipb.SCShenfaAdvanced {
	shenfaAdvanced := &uipb.SCShenfaAdvanced{}
	shenfaAdvanced.AdvancedId = &advancedId
	shenfaAdvanced.ShenfaId = &shenfaId
	shenfaAdvanced.Bless = &bless
	shenfaAdvanced.BlessTime = &blessTime
	typeInt := int32(typ)
	shenfaAdvanced.AdvancedType = &typeInt
	shenfaAdvanced.IsDouble = &isDouble
	shenfaAdvanced.TotalBless = &totalBless
	return shenfaAdvanced
}

func BuildShenfaInfo(info *shenfatypes.ShenfaInfo) *uipb.ShenfaInfo {
	shenfaInfo := &uipb.ShenfaInfo{}
	advancedId := info.AdvanceId
	shenfaInfo.AdvancedId = &advancedId
	shenfaId := info.ShenfaId
	shenfaInfo.ShenfaId = &shenfaId
	unrealLevel := info.UnrealLevel
	unrealPro := info.UnrealPro
	shenfaInfo.UnrealLevel = &unrealLevel
	shenfaInfo.UnrealPro = &unrealPro
	for _, skinInfo := range info.SkinList {
		temp := buildShenFaSkinInfo(skinInfo)
		shenfaInfo.SkinList = append(shenfaInfo.SkinList, temp)
	}
	return shenfaInfo
}

func buildShenFaSkinInfo(info *shenfatypes.ShenfaSkinInfo) *uipb.ShenFaSkinInfo {
	skinInfo := &uipb.ShenFaSkinInfo{}
	shenFaId := info.ShenfaId
	level := info.Level
	pro := info.UpPro

	skinInfo.ShenFaId = &shenFaId
	skinInfo.Level = &level
	skinInfo.Pro = &pro
	return skinInfo
}

func BuildSCShenfaHidden(hiddenFlag bool) *uipb.SCShenfaHidden {
	scShenfaHidden := &uipb.SCShenfaHidden{}
	scShenfaHidden.Hidden = &hiddenFlag
	return scShenfaHidden
}

func BuildSCShenFaUpstar(shenFaId int32, level int32, pro int32) *uipb.SCShenFaUpstar {
	shenFaUpstar := &uipb.SCShenFaUpstar{}
	shenFaUpstar.ShenFaId = &shenFaId
	shenFaUpstar.Level = &level
	shenFaUpstar.UpPro = &pro
	return shenFaUpstar
}
