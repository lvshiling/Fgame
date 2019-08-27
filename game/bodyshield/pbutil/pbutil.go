package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playerbshiled "fgame/fgame/game/bodyshield/player"
	bodyshieldtypes "fgame/fgame/game/bodyshield/types"
	commontypes "fgame/fgame/game/common/types"
)

func BuildSCBodyShieldGet(bShieldInfo *playerbshiled.PlayerBodyShieldObject) *uipb.SCBodyShieldGet {
	bodyShieldGet := &uipb.SCBodyShieldGet{}
	advancedId := int32(bShieldInfo.AdvanceId)
	bodyShieldGet.AdvancedId = &advancedId
	jinjiadanLevel := bShieldInfo.JinjiadanLevel
	progress := bShieldInfo.JinjiadanPro
	bodyShieldGet.JinjiadanLevel = &jinjiadanLevel
	bodyShieldGet.Progress = &progress
	bless := bShieldInfo.Bless
	bodyShieldGet.Bless = &bless
	blessTime := bShieldInfo.BlessTime
	bodyShieldGet.BlessTime = &blessTime
	return bodyShieldGet
}

func BuildSCBodyShieldJJDan(level int32, progress int32) *uipb.SCBodyShieldJJDan {
	bodyShieldJJDan := &uipb.SCBodyShieldJJDan{}
	bodyShieldJJDan.Level = &level
	bodyShieldJJDan.Progress = &progress
	return bodyShieldJJDan
}

func BuildSCBodyShieldAdavancedFinshed(advancedId int32, typ commontypes.AdvancedType) *uipb.SCBodyShieldAdvanced {
	bodyShieldAdvanced := &uipb.SCBodyShieldAdvanced{}
	bodyShieldAdvanced.AdvancedId = &advancedId
	typeInt := int32(typ)
	bodyShieldAdvanced.AdvancedType = &typeInt
	return bodyShieldAdvanced
}

func BuildSCBodyShieldAdavanced(advancedId int32, bless int32, totalBless int32, blessTime int64, typ commontypes.AdvancedType, isDouble bool) *uipb.SCBodyShieldAdvanced {
	bodyShieldAdvanced := &uipb.SCBodyShieldAdvanced{}
	bodyShieldAdvanced.AdvancedId = &advancedId
	bodyShieldAdvanced.Bless = &bless
	bodyShieldAdvanced.BlessTime = &blessTime
	typeInt := int32(typ)
	bodyShieldAdvanced.AdvancedType = &typeInt
	bodyShieldAdvanced.IsDouble = &isDouble
	bodyShieldAdvanced.TotalBless = &totalBless
	return bodyShieldAdvanced
}

func BuildBodyShieldInfo(info *bodyshieldtypes.BodyShieldInfo) *uipb.BodyShieldInfo {
	bodyShieldInfo := &uipb.BodyShieldInfo{}
	jinjiadanLevel := info.JinjiadanLevel
	progress := info.JinjiadanPro
	bodyShieldInfo.JinjiadanLevel = &jinjiadanLevel
	bodyShieldInfo.Progress = &progress
	advanceId := info.AdvancedId
	bodyShieldInfo.AdvancedId = &advanceId

	return bodyShieldInfo
}

func BuildShieldInfo(info *bodyshieldtypes.ShieldInfo) *uipb.ShieldInfo {
	shieldInfo := &uipb.ShieldInfo{}
	shieldId := info.ShieldId
	progress := info.Progress

	shieldInfo.ShieldId = &shieldId
	shieldInfo.Progress = &progress
	return shieldInfo
}

func BuildSCShieldGet(bShieldInfo *playerbshiled.PlayerBodyShieldObject) *uipb.SCShieldGet {
	shieldGet := &uipb.SCShieldGet{}
	shieldId := bShieldInfo.ShieldId
	shieldPro := bShieldInfo.ShieldPro
	shieldGet.ShieldId = &shieldId
	shieldGet.Progress = &shieldPro
	return shieldGet
}

func BuildSCShieldAdvanced(shieldId int32, pro int32, typ commontypes.AdvancedType, isDouble bool) *uipb.SCShieldAdvanced {
	shieldAdvanced := &uipb.SCShieldAdvanced{}
	shieldAdvanced.ShieldId = &shieldId
	shieldAdvanced.Progress = &pro
	typeInt := int32(typ)
	shieldAdvanced.AdvancedType = &typeInt
	shieldAdvanced.IsDouble = &isDouble
	return shieldAdvanced
}
