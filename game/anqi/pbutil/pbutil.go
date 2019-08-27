package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playeranqi "fgame/fgame/game/anqi/player"
	anqitypes "fgame/fgame/game/anqi/types"
	commontypes "fgame/fgame/game/common/types"
)

func BuildSCAnqiGet(bShieldInfo *playeranqi.PlayerAnqiObject) *uipb.SCAnqiGet {
	scAnqiGet := &uipb.SCAnqiGet{}
	advancedId := int32(bShieldInfo.AdvanceId)
	scAnqiGet.AdvancedId = &advancedId
	anqiDanLevel := bShieldInfo.AnqiDanLevel
	progress := bShieldInfo.AnqiDanPro
	scAnqiGet.AnqiDanLevel = &anqiDanLevel
	scAnqiGet.Progress = &progress
	bless := bShieldInfo.Bless
	scAnqiGet.Bless = &bless
	blessTime := bShieldInfo.BlessTime
	scAnqiGet.BlessTime = &blessTime
	return scAnqiGet
}

func BuildSCAnqiEatDan(level int32, progress int32) *uipb.SCAnqiEatDan {
	scAnqiEatDan := &uipb.SCAnqiEatDan{}
	scAnqiEatDan.Level = &level
	scAnqiEatDan.Progress = &progress
	return scAnqiEatDan
}

func BuildSCAnqiAdavancedFinshed(advancedId int32, typ commontypes.AdvancedType) *uipb.SCAnqiAdvanced {
	anqiAdvanced := &uipb.SCAnqiAdvanced{}
	anqiAdvanced.AdvancedId = &advancedId
	typInt := int32(typ)
	anqiAdvanced.AdvancedType = &typInt
	return anqiAdvanced
}

func BuildSCAnqiAdavanced(advancedId int32, bless int32, totalBless int32, blessTime int64, typ commontypes.AdvancedType, isDouble bool) *uipb.SCAnqiAdvanced {
	anqiAdvanced := &uipb.SCAnqiAdvanced{}
	anqiAdvanced.AdvancedId = &advancedId
	anqiAdvanced.Bless = &bless
	anqiAdvanced.BlessTime = &blessTime
	typInt := int32(typ)
	anqiAdvanced.AdvancedType = &typInt
	anqiAdvanced.IsDouble = &isDouble
	anqiAdvanced.TotalBless = &totalBless
	return anqiAdvanced
}

func BuildSCAnqiInfo(info *anqitypes.AnqiInfo) *uipb.SCAnqiInfo {
	scAnqiInfo := &uipb.SCAnqiInfo{}
	scAnqiInfo.AnqiInfo = BuildAnqiInfo(info)
	return scAnqiInfo
}

func BuildAnqiInfo(info *anqitypes.AnqiInfo) *uipb.AnqiInfo {
	anqiInfo := &uipb.AnqiInfo{}
	anqiDanLevel := info.AnqiDanLevel
	progress := info.AnqiDanPro
	anqiInfo.AnqiDanLevel = &anqiDanLevel
	anqiInfo.Progress = &progress
	advanceId := info.AdvancedId
	anqiInfo.AdvancedId = &advanceId

	return anqiInfo
}
