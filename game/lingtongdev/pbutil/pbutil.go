package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	commontypes "fgame/fgame/game/common/types"
	lingtongdevcommon "fgame/fgame/game/lingtongdev/common"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
)

func BuildSCLingTongDevGet(classType int32, lingTongDevObj *playerlingtongdev.PlayerLingTongDevObject, container *playerlingtongdev.LingTongOtherContainer) *uipb.SCLingTongDevGet {
	lingTongDevGet := &uipb.SCLingTongDevGet{}
	lingTongDevGet.ClassType = &classType

	lingTongDevGet.LingTongInfo = buildLingTongDev(lingTongDevObj)

	if container != nil {
		for _, typeOtherMap := range container.GetOtherMap() {
			for _, otherObj := range typeOtherMap {
				lingTongDevGet.LingTongSkinList = append(lingTongDevGet.LingTongSkinList, buildLingTongDevOther(otherObj))
			}
		}
	}
	return lingTongDevGet
}

func buildLingTongDev(obj *playerlingtongdev.PlayerLingTongDevObject) *uipb.LingTongDevInfo {
	lingTongDevInfo := &uipb.LingTongDevInfo{}
	advancedId := obj.GetAdvancedId()
	seqId := obj.GetSeqId()
	unrealLevel := obj.GetUnrealLevel()
	unrealPro := obj.GetUnrealPro()
	bless := obj.GetBless()
	blessTime := obj.GetBlessTime()
	culLevel := obj.GetCulLevel()
	culPro := obj.GetCulPro()
	tonglingLevel := obj.GetTongLingLevel()
	tonglingPro := obj.GetTongLingPro()
	hidden := false
	if obj.GetHidden() == 1 {
		hidden = true
	}

	for _, unrealId := range obj.GetUnrealList() {
		lingTongDevInfo.UnrealList = append(lingTongDevInfo.UnrealList, int32(unrealId))
	}

	lingTongDevInfo.AdvancedId = &advancedId
	lingTongDevInfo.SeqId = &seqId
	lingTongDevInfo.UnrealLevel = &unrealLevel
	lingTongDevInfo.UnrealPro = &unrealPro
	lingTongDevInfo.Bless = &bless
	lingTongDevInfo.BlessTime = &blessTime
	lingTongDevInfo.CulLevel = &culLevel
	lingTongDevInfo.CulPro = &culPro
	lingTongDevInfo.TongLingLevel = &tonglingLevel
	lingTongDevInfo.TongLingPro = &tonglingPro
	lingTongDevInfo.Hidden = &hidden
	return lingTongDevInfo
}

func buildLingTongDevOther(obj *playerlingtongdev.PlayerLingTongOtherObject) *uipb.LingTongDevSkinInfo {
	lingTongDevSkinInfo := &uipb.LingTongDevSkinInfo{}
	seqId := obj.GetSeqId()
	level := obj.GetLevel()
	pro := obj.GetUpPro()

	lingTongDevSkinInfo.SeqId = &seqId
	lingTongDevSkinInfo.Level = &level
	lingTongDevSkinInfo.UpPro = &pro
	return lingTongDevSkinInfo
}

func BuildSCLingTongDevUnrealDan(classType int32, level int32, progress int32) *uipb.SCLingTongDevUnrealDan {
	lingTongDevUnrealDan := &uipb.SCLingTongDevUnrealDan{}
	lingTongDevUnrealDan.ClassType = &classType
	lingTongDevUnrealDan.UnrealLevel = &level
	lingTongDevUnrealDan.UnrealPro = &progress
	return lingTongDevUnrealDan
}

func BuildSCLingTongDevCulDan(classType int32, level int32, progress int32) *uipb.SCLingTongDevCulDan {
	lingTongDevCulDan := &uipb.SCLingTongDevCulDan{}
	lingTongDevCulDan.CulLevel = &level
	lingTongDevCulDan.ClassType = &classType
	lingTongDevCulDan.CulPro = &progress
	return lingTongDevCulDan
}

func BuildSCLingTongDevUnreal(classType int32, seqId int32) *uipb.SCLingTongDevUnreal {
	lingTongDevUnreal := &uipb.SCLingTongDevUnreal{}
	lingTongDevUnreal.ClassType = &classType
	lingTongDevUnreal.SeqId = &seqId
	return lingTongDevUnreal
}

func BuildSCLingTongDevUnload(classType int32, seqId int32) *uipb.SCLingTongDevUnload {
	unLoad := &uipb.SCLingTongDevUnload{}
	unLoad.ClassType = &classType
	unLoad.SeqId = &seqId
	return unLoad
}

func BuildSCLingTongDevAdavancedFinshed(classType int32, advancedId int32, seqId int32, typ commontypes.AdvancedType) *uipb.SCLingTongDevAdvanced {
	lingTongDevAdvanced := &uipb.SCLingTongDevAdvanced{}
	lingTongDevAdvanced.AdvancedId = &advancedId
	lingTongDevAdvanced.ClassType = &classType
	lingTongDevAdvanced.SeqId = &seqId
	typInt := int32(typ)
	lingTongDevAdvanced.AdvancedType = &typInt
	return lingTongDevAdvanced
}

func BuildSCLingTongDevAdavanced(classType int32, advancedId int32, seqId int32, bless int32, totalBless int32, blessTime int64, typ commontypes.AdvancedType, isDouble bool) *uipb.SCLingTongDevAdvanced {
	lingTongDevAdvanced := &uipb.SCLingTongDevAdvanced{}
	lingTongDevAdvanced.ClassType = &classType
	lingTongDevAdvanced.AdvancedId = &advancedId
	lingTongDevAdvanced.SeqId = &seqId
	lingTongDevAdvanced.Bless = &bless
	lingTongDevAdvanced.BlessTime = &blessTime
	typInt := int32(typ)
	lingTongDevAdvanced.AdvancedType = &typInt
	lingTongDevAdvanced.IsDouble = &isDouble
	lingTongDevAdvanced.TotalBless = &totalBless
	return lingTongDevAdvanced
}

func BuildSCLingTongDevHidden(classType int32, hiddenFlag bool) *uipb.SCLingTongDevHidden {
	lingTongDevHidden := &uipb.SCLingTongDevHidden{}
	lingTongDevHidden.ClassType = &classType
	lingTongDevHidden.Hidden = &hiddenFlag
	return lingTongDevHidden
}

func BuildSCLingTongDevTongLing(classType int32, level int32, pro int32) *uipb.SCLingTongDevTongLing {
	scLingTongDevTongLing := &uipb.SCLingTongDevTongLing{}
	scLingTongDevTongLing.ClassType = &classType
	scLingTongDevTongLing.TongLingLevel = &level
	scLingTongDevTongLing.TongLingPro = &pro
	return scLingTongDevTongLing
}

func BuildSCLingTongDevUpstar(classType int32, seqId int32, level int32, pro int32) *uipb.SCLingTongDevUpstar {
	lingTongDevUpstar := &uipb.SCLingTongDevUpstar{}
	lingTongDevUpstar.ClassType = &classType
	lingTongDevUpstar.SeqId = &seqId
	lingTongDevUpstar.Level = &level
	lingTongDevUpstar.UpPro = &pro
	return lingTongDevUpstar
}

func BuildAllLingTongDevInfo(info *lingtongdevcommon.AllLingTongDevInfo) *uipb.AllLingTongDevInfo {
	allLingTongDevInfo := &uipb.AllLingTongDevInfo{}
	for _, tempInfo := range info.LingTongDevList {
		allLingTongDevInfo.LingTongDevList = append(allLingTongDevInfo.LingTongDevList, buildLingTongDevCacheInfo(tempInfo))
	}
	return allLingTongDevInfo
}

func buildLingTongDevCacheInfo(info *lingtongdevcommon.LingTongDevInfo) *uipb.LingTongDevCacheInfo {
	lingTongDevCacheInfo := &uipb.LingTongDevCacheInfo{}
	classType := info.ClassType
	advancedId := int32(info.AdvanceId)
	seqId := info.SeqId
	unrealLevel := info.UnrealLevel
	unrealPro := info.UnrealPro
	culLevel := info.CulLevel
	culPro := info.CulPro
	tongLingLevel := info.TongLingLevel
	tongLingPro := info.TongLingPro

	lingTongDevCacheInfo.ClassType = &classType
	lingTongDevCacheInfo.AdvanceId = &advancedId
	lingTongDevCacheInfo.SeqId = &seqId
	lingTongDevCacheInfo.UnrealLevel = &unrealLevel
	lingTongDevCacheInfo.UnrealPro = &unrealPro
	lingTongDevCacheInfo.CulLevel = &culLevel
	lingTongDevCacheInfo.CulPro = &culPro
	lingTongDevCacheInfo.TongLingLevel = &tongLingLevel
	lingTongDevCacheInfo.TongLingPro = &tongLingPro

	for _, skinInfo := range info.SkinList {
		temp := buildLingTongDevSkinInfo(skinInfo)
		lingTongDevCacheInfo.SkinList = append(lingTongDevCacheInfo.SkinList, temp)
	}

	return lingTongDevCacheInfo
}

func buildLingTongDevSkinInfo(info *lingtongdevcommon.LingTongDevSkinInfo) *uipb.LingTongDevSkinInfo {
	skinInfo := &uipb.LingTongDevSkinInfo{}
	seqId := info.SeqId
	level := info.Level
	pro := info.UpPro

	skinInfo.SeqId = &seqId
	skinInfo.Level = &level
	skinInfo.UpPro = &pro
	return skinInfo
}
