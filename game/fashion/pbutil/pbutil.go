package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playerfashion "fgame/fgame/game/fashion/player"
	fashiontypes "fgame/fgame/game/fashion/types"
)

func BuildSCFashionGet(fashionWear int32, fashionMap map[fashiontypes.FashionType]map[int32]*playerfashion.PlayerFashionObject, trialObjMap map[int32]*playerfashion.PlayerFashionTrialObject) *uipb.SCFashionGet {
	fashionGet := &uipb.SCFashionGet{}
	fashionGet.FashionWear = &fashionWear
	for _, fashionTypeMap := range fashionMap {
		for _, fashionObject := range fashionTypeMap {
			fashionGet.FashionList = append(fashionGet.FashionList, buildFashion(fashionObject))
		}
	}

	for _, obj := range trialObjMap {
		trialId := obj.GetTrialFashionId()
		expireTime := obj.GetExpireTime()
		fashionGet.TrialList = append(fashionGet.TrialList, buildTrialInfo(trialId, expireTime))

	}

	return fashionGet
}

func BuildSCFashionActive(fashionId int32, activeTime int64) *uipb.SCFashionActive {
	fashionActive := &uipb.SCFashionActive{}
	fashionActive.FashionId = &fashionId
	fashionActive.ActiveTime = &activeTime
	return fashionActive
}

func BuildSCFashionWear(fashionId int32) *uipb.SCFashionWear {
	fashionWear := &uipb.SCFashionWear{}
	fashionWear.FashionId = &fashionId
	return fashionWear
}

func BuildSCFashionUnload(fashionWear int32) *uipb.SCFashionUnLoad {
	fashionUnLoad := &uipb.SCFashionUnLoad{}
	fashionUnLoad.FashionWear = &fashionWear
	return fashionUnLoad
}

func BuildSCFashionUpstar(fashionId int32, star int32, pro int32, sucess bool) *uipb.SCFahionUpstar {
	fahionUpstar := &uipb.SCFahionUpstar{}
	fahionUpstar.FashionId = &fashionId
	fahionUpstar.Level = &star
	fahionUpstar.UpPro = &pro
	fahionUpstar.Result = &sucess
	return fahionUpstar
}

func BuildSCFashionRemove(fashionId int32) *uipb.SCFashionRemove {
	fashionRemove := &uipb.SCFashionRemove{}
	fashionRemove.FashionId = &fashionId
	return fashionRemove
}

func BuildSCFashionTrialNotice(trialId int32, expireTime int64) *uipb.SCFashionTrialNotice {
	scFashionTrialNotice := &uipb.SCFashionTrialNotice{}
	scFashionTrialNotice.TrialFashionId = &trialId
	scFashionTrialNotice.ExpireTime = &expireTime
	return scFashionTrialNotice
}

func BuildSCFashionTrialOverdueNotice(trialId int32, overdueType fashiontypes.FashionTrialOverdueType) *uipb.SCFashionTrialOverdueNotice {
	scFashionTrialOverdueNotice := &uipb.SCFashionTrialOverdueNotice{}
	scFashionTrialOverdueNotice.TrialFashionId = &trialId
	typInt := int32(overdueType)
	scFashionTrialOverdueNotice.OverdueType = &typInt
	return scFashionTrialOverdueNotice
}

func buildFashion(fashionObj *playerfashion.PlayerFashionObject) *uipb.FashionInfo {
	fashionInfo := &uipb.FashionInfo{}
	fashionId := fashionObj.FashionId
	level := fashionObj.Star
	pro := fashionObj.UpStarPro
	activeTime := fashionObj.ActiveTime
	fashionInfo.FashionId = &fashionId
	fashionInfo.Level = &level
	fashionInfo.UpPro = &pro
	fashionInfo.AcitveTime = &activeTime
	return fashionInfo
}
func buildTrialInfo(trialId int32, expireTime int64) *uipb.TrialInfo {
	trialInfo := &uipb.TrialInfo{}
	trialInfo.TrialFashionId = &trialId
	trialInfo.ExpireTime = &expireTime

	return trialInfo
}
