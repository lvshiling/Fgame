package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	commontypes "fgame/fgame/game/common/types"
	playershihunfan "fgame/fgame/game/shihunfan/player"
)

func BuildSCShiHunFanGet(info *playershihunfan.PlayerShiHunFanObject) *uipb.SCShihunfanGet {
	scShiHunFanGet := &uipb.SCShihunfanGet{}
	chargeVal := info.ChargeVal
	scShiHunFanGet.ShihunfanInfo = buildShiHunFanInfo(info)
	scShiHunFanGet.ShihunfandanInfo = buildShiHunFanDanInfo(info)
	scShiHunFanGet.ChargeVal = &chargeVal
	return scShiHunFanGet
}

func BuildSCShiHunFanAdavancedFinshed(info *playershihunfan.PlayerShiHunFanObject, typ commontypes.AdvancedType) *uipb.SCShihunfanAdvanced {
	shiHunFanAdvanced := &uipb.SCShihunfanAdvanced{}
	chargeVal := info.ChargeVal
	shiHunFanAdvanced.ShihunfanInfo = buildShiHunFanInfo(info)
	shiHunFanAdvanced.ChargeVal = &chargeVal
	typInt := int32(typ)
	shiHunFanAdvanced.AdvancedType = &typInt
	return shiHunFanAdvanced
}

func BuildSCShiHunFanAdavanced(info *playershihunfan.PlayerShiHunFanObject, typ commontypes.AdvancedType, isDouble bool, isAutoBuy bool) *uipb.SCShihunfanAdvanced {
	shiHunFanAdvanced := &uipb.SCShihunfanAdvanced{}
	chargeVal := info.ChargeVal
	shiHunFanAdvanced.ShihunfanInfo = buildShiHunFanInfo(info)
	shiHunFanAdvanced.ChargeVal = &chargeVal
	typInt := int32(typ)
	shiHunFanAdvanced.AdvancedType = &typInt
	shiHunFanAdvanced.IsDouble = &isDouble
	shiHunFanAdvanced.BuyFlag = &isAutoBuy
	return shiHunFanAdvanced
}

func BuildSCShiHunFanEatDan(info *playershihunfan.PlayerShiHunFanObject) *uipb.SCShihunfanDanAdvanced {
	scshiHunFanEatDan := &uipb.SCShihunfanDanAdvanced{}
	scshiHunFanEatDan.ShihunfandanInfo = buildShiHunFanDanInfo(info)
	return scshiHunFanEatDan
}

func BuildSCShiHunFanChargeVary(num int32) *uipb.SCShihunfanChargeVary {
	scMsg := &uipb.SCShihunfanChargeVary{}
	scMsg.ChargeVal = &num

	return scMsg
}

func buildShiHunFanInfo(info *playershihunfan.PlayerShiHunFanObject) *uipb.ShihunfanInfo {
	shiHunFanInfo := &uipb.ShihunfanInfo{}

	advanceId := int32(info.AdvanceId)
	bless := info.Bless
	blessTime := info.BlessTime

	shiHunFanInfo.AdvancedId = &advanceId
	shiHunFanInfo.Bless = &bless
	shiHunFanInfo.BlessTime = &blessTime
	return shiHunFanInfo
}

func buildShiHunFanDanInfo(info *playershihunfan.PlayerShiHunFanObject) *uipb.ShihunfanDanInfo {
	danInfo := &uipb.ShihunfanDanInfo{}

	danLevel := info.DanLevel
	danPro := info.DanPro

	danInfo.Level = &danLevel
	danInfo.Progress = &danPro
	return danInfo
}
