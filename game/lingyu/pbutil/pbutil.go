package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	commontypes "fgame/fgame/game/common/types"
	playerlingyu "fgame/fgame/game/lingyu/player"
	lingyutypes "fgame/fgame/game/lingyu/types"
)

func BuildSCLingyuGet(lingyuInfo *playerlingyu.PlayerLingyuObject, lingYuOtherMap map[lingyutypes.LingyuType]map[int32]*playerlingyu.PlayerLingyuOtherObject) *uipb.SCLingyuGet {
	lingyuGet := &uipb.SCLingyuGet{}
	advancedId := int32(lingyuInfo.AdvanceId)
	lingyuGet.AdvancedId = &advancedId
	lingyuId := lingyuInfo.LingyuId
	lingyuGet.LingyuId = &lingyuId
	unrealLevel := lingyuInfo.UnrealLevel
	unrealPro := lingyuInfo.UnrealPro
	lingyuGet.UnrealLevel = &unrealLevel
	lingyuGet.UnrealPro = &unrealPro
	bless := lingyuInfo.Bless
	lingyuGet.Bless = &bless
	blessTime := lingyuInfo.BlessTime
	lingyuGet.BlessTime = &blessTime

	hidden := false
	if lingyuInfo.Hidden == 1 {
		hidden = true
	}
	lingyuGet.Hidden = &hidden
	for _, unrealId := range lingyuInfo.UnrealList {
		lingyuGet.UnrealList = append(lingyuGet.UnrealList, int32(unrealId))
	}

	for _, lingYuTypeOtherMap := range lingYuOtherMap {
		for _, lingYuOtherObj := range lingYuTypeOtherMap {
			lingyuGet.LingYuSkinList = append(lingyuGet.LingYuSkinList, buildLingYuOther(lingYuOtherObj))
		}
	}
	return lingyuGet
}

func buildLingYuOther(lingYuOtherObj *playerlingyu.PlayerLingyuOtherObject) *uipb.LingYuSkinInfo {
	lingYuSkinInfo := &uipb.LingYuSkinInfo{}
	lingYuId := lingYuOtherObj.LingYuId
	level := lingYuOtherObj.Level
	pro := lingYuOtherObj.UpPro

	lingYuSkinInfo.LingYuId = &lingYuId
	lingYuSkinInfo.Level = &level
	lingYuSkinInfo.Pro = &pro
	return lingYuSkinInfo
}

func BuildSCLingyuUnrealDan(level int32, progress int32) *uipb.SCLingyuUnrealDan {
	lingyuUnrealDan := &uipb.SCLingyuUnrealDan{}
	lingyuUnrealDan.Level = &level
	lingyuUnrealDan.Progress = &progress
	return lingyuUnrealDan
}

func BuildSCLingyuUnreal(lingyuId int32) *uipb.SCLingyuUnreal {
	lingyuUnreal := &uipb.SCLingyuUnreal{}
	lingyuUnreal.LingyuId = &lingyuId

	return lingyuUnreal
}

func BuildSCLingyuUnload(lingyuId int32) *uipb.SCLingyuUnload {
	unLoad := &uipb.SCLingyuUnload{}
	unLoad.LingyuId = &lingyuId
	return unLoad
}

func BuildSCLingyuAdavancedFinshed(advancedId int32, lingyuId int32, typ commontypes.AdvancedType) *uipb.SCLingyuAdvanced {
	lingyuAdvanced := &uipb.SCLingyuAdvanced{}
	lingyuAdvanced.AdvancedId = &advancedId
	lingyuAdvanced.LingyuId = &lingyuId
	typeInt := int32(typ)
	lingyuAdvanced.AdvancedType = &typeInt
	return lingyuAdvanced
}

func BuildSCLingyuAdavanced(advancedId int32, lingyuId int32, bless int32, totalBless int32, blessTime int64, typ commontypes.AdvancedType, isDouble bool) *uipb.SCLingyuAdvanced {
	lingyuAdvanced := &uipb.SCLingyuAdvanced{}
	lingyuAdvanced.AdvancedId = &advancedId
	lingyuAdvanced.LingyuId = &lingyuId
	lingyuAdvanced.Bless = &bless
	lingyuAdvanced.BlessTime = &blessTime
	typeInt := int32(typ)
	lingyuAdvanced.AdvancedType = &typeInt
	lingyuAdvanced.IsDouble = &isDouble
	lingyuAdvanced.TotalBless = &totalBless
	return lingyuAdvanced
}

func BuildLingyuInfo(info *lingyutypes.LingyuInfo) *uipb.LingyuInfo {
	lingyuInfo := &uipb.LingyuInfo{}
	advancedId := info.AdvanceId
	lingyuInfo.AdvancedId = &advancedId
	lingyuId := info.LingyuId
	lingyuInfo.LingyuId = &lingyuId
	unrealLevel := info.UnrealLevel
	unrealPro := info.UnrealPro
	lingyuInfo.UnrealLevel = &unrealLevel
	lingyuInfo.UnrealPro = &unrealPro
	for _, skinInfo := range info.SkinList {
		temp := buildLingYuSkinInfo(skinInfo)
		lingyuInfo.SkinList = append(lingyuInfo.SkinList, temp)
	}
	return lingyuInfo
}

func buildLingYuSkinInfo(info *lingyutypes.LingyuSkinInfo) *uipb.LingYuSkinInfo {
	skinInfo := &uipb.LingYuSkinInfo{}
	lingYuId := info.LingyuId
	level := info.Level
	pro := info.UpPro

	skinInfo.LingYuId = &lingYuId
	skinInfo.Level = &level
	skinInfo.Pro = &pro
	return skinInfo
}

func BuildSCLingyuHidden(hiddenFlag bool) *uipb.SCLingyuHidden {
	lingyuHidden := &uipb.SCLingyuHidden{}
	lingyuHidden.Hidden = &hiddenFlag
	return lingyuHidden
}

func BuildSCLingYuUpstar(lingYuId int32, level int32, pro int32) *uipb.SCLingYuUpstar {
	lingYuUpstar := &uipb.SCLingYuUpstar{}
	lingYuUpstar.LingYuId = &lingYuId
	lingYuUpstar.Level = &level
	lingYuUpstar.UpPro = &pro
	return lingYuUpstar
}
