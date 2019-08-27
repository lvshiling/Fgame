package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	commontypes "fgame/fgame/game/common/types"
	wingcommon "fgame/fgame/game/wing/common"
	playerwing "fgame/fgame/game/wing/player"
	wingtypes "fgame/fgame/game/wing/types"
)

func BuildSCWingGet(wingInfo *playerwing.PlayerWingObject, wingTrialInfo *playerwing.PlayerWingTrialObject, wingOtherMap map[wingtypes.WingType]map[int32]*playerwing.PlayerWingOtherObject) *uipb.SCWingGet {
	wingGet := &uipb.SCWingGet{}
	advancedId := int32(wingInfo.AdvanceId)
	wingGet.AdvancedId = &advancedId
	wingId := wingInfo.WingId
	wingGet.WingId = &wingId
	unrealLevel := wingInfo.UnrealLevel
	unrealPro := wingInfo.UnrealPro
	wingGet.UnrealLevel = &unrealLevel
	wingGet.UnrealPro = &unrealPro
	bless := wingInfo.Bless
	wingGet.Bless = &bless
	blessTime := wingInfo.BlessTime
	wingGet.BlessTime = &blessTime

	if wingTrialInfo.TrialOrderId != 0 {
		trialOrderId := wingTrialInfo.TrialOrderId
		activeTime := wingTrialInfo.ActiveTime
		wingGet.TrialOrderId = &trialOrderId
		wingGet.ActiveTime = &activeTime
	}

	hidden := false
	if wingInfo.Hidden == 1 {
		hidden = true
	}
	wingGet.Hidden = &hidden
	for _, unrealId := range wingInfo.UnrealList {
		wingGet.UnrealList = append(wingGet.UnrealList, int32(unrealId))
	}
	for _, wingTypeOtherMap := range wingOtherMap {
		for _, wingOtherObj := range wingTypeOtherMap {

			wingGet.WingSkinList = append(wingGet.WingSkinList, buildWingOther(wingOtherObj))
		}
	}

	wingGet.FeatherInfo = buildFeather(wingInfo)
	return wingGet
}

func buildFeather(wingInfo *playerwing.PlayerWingObject) *uipb.FeatherInfo {
	featherInfo := &uipb.FeatherInfo{}
	featherId := wingInfo.FeatherId
	progress := wingInfo.FeatherPro

	featherInfo.FeatherId = &featherId
	featherInfo.Progress = &progress
	return featherInfo
}

func buildWingOther(wingOtherObj *playerwing.PlayerWingOtherObject) *uipb.WingSkinInfo {
	wingSkinInfo := &uipb.WingSkinInfo{}
	wingId := wingOtherObj.WingId
	level := wingOtherObj.Level
	pro := wingOtherObj.UpPro

	wingSkinInfo.WingId = &wingId
	wingSkinInfo.Level = &level
	wingSkinInfo.Pro = &pro
	return wingSkinInfo
}

func BuildSCWingUnrealDan(level int32, progress int32) *uipb.SCWingUnrealDan {
	wingUnrealDan := &uipb.SCWingUnrealDan{}
	wingUnrealDan.Level = &level
	wingUnrealDan.Progress = &progress
	return wingUnrealDan
}

func BuildSCWingUnreal(wingId int32) *uipb.SCWingUnreal {
	wingUnreal := &uipb.SCWingUnreal{}
	wingUnreal.WingId = &wingId
	return wingUnreal
}

func BuildSCWingUnload(wingId int32) *uipb.SCWingUnload {
	unLoad := &uipb.SCWingUnload{}
	unLoad.WingId = &wingId
	return unLoad
}

func BuildSCWingAdavancedFinshed(advancedId int32, wingId int32, typ commontypes.AdvancedType) *uipb.SCWingAdvanced {
	wingAdvanced := &uipb.SCWingAdvanced{}
	wingAdvanced.AdvancedId = &advancedId
	wingAdvanced.WingId = &wingId
	typInt := int32(typ)
	wingAdvanced.AdvancedType = &typInt
	return wingAdvanced
}

func BuildSCWingAdavanced(advancedId int32, wingId int32, bless int32, totalBless int32, blessTime int64, typ commontypes.AdvancedType, isDouble bool) *uipb.SCWingAdvanced {
	wingAdvanced := &uipb.SCWingAdvanced{}
	wingAdvanced.AdvancedId = &advancedId
	wingAdvanced.WingId = &wingId
	wingAdvanced.Bless = &bless
	wingAdvanced.BlessTime = &blessTime
	typInt := int32(typ)
	wingAdvanced.AdvancedType = &typInt
	wingAdvanced.IsDouble = &isDouble
	wingAdvanced.TotalBless = &totalBless
	return wingAdvanced
}

func BuildWingInfo(info *wingcommon.WingInfo) *uipb.WingInfo {
	wingInfo := &uipb.WingInfo{}
	advancedId := info.AdvanceId
	wingInfo.AdvancedId = &advancedId
	wingId := info.WingId
	wingInfo.WingId = &wingId
	unrealLevel := info.UnrealLevel
	unrealPro := info.UnrealPro
	wingInfo.UnrealLevel = &unrealLevel
	wingInfo.UnrealPro = &unrealPro
	for _, skinInfo := range info.SkinList {
		temp := buildWingSkinInfo(skinInfo)
		wingInfo.SkinList = append(wingInfo.SkinList, temp)
	}
	return wingInfo
}

func buildWingSkinInfo(info *wingcommon.WingSkinInfo) *uipb.WingSkinInfo {
	wingSkinInfo := &uipb.WingSkinInfo{}
	wingId := info.WingId
	level := info.Level
	pro := info.UpPro

	wingSkinInfo.WingId = &wingId
	wingSkinInfo.Level = &level
	wingSkinInfo.Pro = &pro
	return wingSkinInfo
}

func BuildFeatherInfo(info *wingcommon.FeatherInfo) *uipb.FeatherInfo {
	featherInfo := &uipb.FeatherInfo{}
	featherId := info.FeatherId
	progress := info.Progress

	featherInfo.FeatherId = &featherId
	featherInfo.Progress = &progress
	return featherInfo
}

func BuildSCFeatherGet(wingInfo *playerwing.PlayerWingObject) *uipb.SCFeatherGet {
	featherGet := &uipb.SCFeatherGet{}
	featherId := wingInfo.FeatherId
	featherPro := wingInfo.FeatherPro

	featherGet.FeatherId = &featherId
	featherGet.Progress = &featherPro
	return featherGet
}

func BuildSCFeatherAdvanced(featherId int32, pro int32, typ commontypes.AdvancedType) *uipb.SCFeatherAdvanced {
	featherAdvanced := &uipb.SCFeatherAdvanced{}
	featherAdvanced.FeatherId = &featherId
	featherAdvanced.Progress = &pro
	typInt := int32(typ)
	featherAdvanced.AdvancedType = &typInt
	return featherAdvanced
}

func BuildSCWingHidden(hiddenFlag bool) *uipb.SCWingHidden {
	wingHidden := &uipb.SCWingHidden{}
	wingHidden.Hidden = &hiddenFlag
	return wingHidden
}

func BuildSCWingTrialCard(trialWingId int32, activeTime int64) *uipb.SCWingTrialCard {
	wingTrialCard := &uipb.SCWingTrialCard{}
	wingTrialCard.TrialWingId = &trialWingId
	wingTrialCard.ActiveTime = &activeTime
	return wingTrialCard
}

func BuildSCWingTrialOverdue(trialWingId int32, bResult bool) *uipb.SCWingTrialOverdue {
	wingTrialOverdue := &uipb.SCWingTrialOverdue{}
	wingTrialOverdue.TrialWingId = &trialWingId
	wingTrialOverdue.BResult = &bResult
	return wingTrialOverdue
}

func BuildSCWingUpstar(wingId int32, level int32, pro int32) *uipb.SCWingUpstar {
	wingUpstar := &uipb.SCWingUpstar{}
	wingUpstar.WingId = &wingId
	wingUpstar.Level = &level
	wingUpstar.UpPro = &pro
	return wingUpstar
}
