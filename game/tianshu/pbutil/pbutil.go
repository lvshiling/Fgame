package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playertianshu "fgame/fgame/game/tianshu/player"
	tianshutypes "fgame/fgame/game/tianshu/types"
)

func BuildSCTianShuActivate(typ tianshutypes.TianShuType) *uipb.SCTianShuActivate {
	scMsg := &uipb.SCTianShuActivate{}
	typInt := int32(typ)
	scMsg.Type = &typInt
	return scMsg
}

func BuildSCTianShuUplevel(isSuccess bool, typ tianshutypes.TianShuType, level int32) *uipb.SCTianShuUplevel {
	scMsg := &uipb.SCTianShuUplevel{}
	scMsg.IsSuccess = &isSuccess
	typInt := int32(typ)
	scMsg.Type = &typInt
	scMsg.Level = &level
	return scMsg
}

func BuildSCTianShuGiftReceive(typ tianshutypes.TianShuType) *uipb.SCTianShuGiftReceive {
	scMsg := &uipb.SCTianShuGiftReceive{}
	typInt := int32(typ)
	scMsg.Type = &typInt

	return scMsg
}

func BuildSCTianShuInfoList(infoList map[tianshutypes.TianShuType]*playertianshu.PlayerTianShuObject, totalChargeNum int32) *uipb.SCTianShuInfoList {
	scMsg := &uipb.SCTianShuInfoList{}
	for typ, obj := range infoList {
		isReceive := obj.GetIsReceive()
		level := obj.GetLevel()
		scMsg.InfoList = append(scMsg.InfoList, buildTianShuInfo(typ, level, isReceive))
	}
	scMsg.TotalChargeNum = &totalChargeNum

	return scMsg
}

func buildTianShuInfo(typ tianshutypes.TianShuType, level, isReceive int32) *uipb.TianShuInfo {
	info := &uipb.TianShuInfo{}
	typInt := int32(typ)
	info.Type = &typInt
	info.IsReceive = &isReceive
	info.Level = &level

	return info
}
