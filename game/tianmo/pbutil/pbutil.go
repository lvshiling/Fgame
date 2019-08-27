package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	commontypes "fgame/fgame/game/common/types"
	playertianmo "fgame/fgame/game/tianmo/player"
	tianmotypes "fgame/fgame/game/tianmo/types"
)

func BuildSCTianMoGet(tianMoInfo *playertianmo.PlayerTianMoObject) *uipb.SCTianMoGet {
	scTianMoGet := &uipb.SCTianMoGet{}
	advancedId := int32(tianMoInfo.AdvanceId)
	scTianMoGet.AdvancedId = &advancedId
	tianMoDanLevel := tianMoInfo.TianMoDanLevel
	progress := tianMoInfo.TianMoDanPro
	scTianMoGet.TianMoDanLevel = &tianMoDanLevel
	scTianMoGet.Progress = &progress
	bless := tianMoInfo.Bless
	scTianMoGet.Bless = &bless
	blessTime := tianMoInfo.BlessTime
	scTianMoGet.BlessTime = &blessTime
	chargeNum := tianMoInfo.ChargeVal
	scTianMoGet.ChargeGold = &chargeNum
	return scTianMoGet
}

func BuildSCTianMoEatDan(level int32, progress int32) *uipb.SCTianMoEatDan {
	scTianMoEatDan := &uipb.SCTianMoEatDan{}
	scTianMoEatDan.Level = &level
	scTianMoEatDan.Progress = &progress
	return scTianMoEatDan
}

func BuildSCTianMoAdavancedFinshed(advancedId int32, typ commontypes.AdvancedType) *uipb.SCTianMoAdvanced {
	tianMoAdvanced := &uipb.SCTianMoAdvanced{}
	tianMoAdvanced.AdvancedId = &advancedId
	typInt := int32(typ)
	tianMoAdvanced.AdvancedType = &typInt
	return tianMoAdvanced
}

func BuildSCTianMoAdavanced(advancedId int32, bless int32, totalBless int32, blessTime int64, typ commontypes.AdvancedType, isDouble bool) *uipb.SCTianMoAdvanced {
	tianMoAdvanced := &uipb.SCTianMoAdvanced{}
	tianMoAdvanced.AdvancedId = &advancedId
	tianMoAdvanced.Bless = &bless
	tianMoAdvanced.BlessTime = &blessTime
	typInt := int32(typ)
	tianMoAdvanced.AdvancedType = &typInt
	tianMoAdvanced.IsDouble = &isDouble
	tianMoAdvanced.TotalBless = &totalBless
	return tianMoAdvanced
}

func BuildSCTianMoInfo(info *tianmotypes.TianMoInfo) *uipb.SCTianMoInfo {
	scTianMoInfo := &uipb.SCTianMoInfo{}
	scTianMoInfo.TianMoInfo = BuildTianMoInfo(info)
	return scTianMoInfo
}

func BuildTianMoInfo(info *tianmotypes.TianMoInfo) *uipb.TianMoInfo {
	tianMoInfo := &uipb.TianMoInfo{}
	tianMoDanLevel := info.TianMoDanLevel
	progress := info.TianMoDanPro
	tianMoInfo.TianMoDanLevel = &tianMoDanLevel
	tianMoInfo.Progress = &progress
	advanceId := info.AdvancedId
	tianMoInfo.AdvancedId = &advanceId

	return tianMoInfo
}

func BuildSCTianMoChargeGold(chargeNum int64) *uipb.SCTianMoChargeGold {
	scMsg := &uipb.SCTianMoChargeGold{}
	scMsg.ChargeGold = &chargeNum

	return scMsg
}
