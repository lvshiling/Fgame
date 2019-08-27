package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
)

func BuildSCLingTongFashionGet(lingTongId int32, fashionMap map[int32]*playerlingtong.PlayerLingTongFashionInfoObject) *uipb.SCLingTongFashionGet {
	lingTongFashionGet := &uipb.SCLingTongFashionGet{}
	lingTongFashionGet.LingTongId = &lingTongId
	for _, fashionObj := range fashionMap {
		lingTongFashionGet.FashionList = append(lingTongFashionGet.FashionList, buildLingTongFashion(fashionObj))
	}
	return lingTongFashionGet
}

func buildLingTongFashion(obj *playerlingtong.PlayerLingTongFashionInfoObject) *uipb.LingTongFashionInfo {
	lingTongFashionInfo := &uipb.LingTongFashionInfo{}
	fashionId := obj.GetFashionId()
	level := obj.GetUpgradeLevel()
	levelPro := obj.GetUpgradePro()
	activateTime := int64(0)
	lingTongFashionTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(fashionId)
	if lingTongFashionTemplate != nil && lingTongFashionTemplate.Time != 0 {
		activateTime = obj.GetActivateTime()
	}
	lingTongFashionInfo.FashionId = &fashionId
	lingTongFashionInfo.Level = &level
	lingTongFashionInfo.LevelPro = &levelPro
	lingTongFashionInfo.ActivteTime = &activateTime
	return lingTongFashionInfo
}

func BuildSCLingTongFashionWear(fashionId int32) *uipb.SCLingTongFashionWear {
	scLingTongFashionWear := &uipb.SCLingTongFashionWear{}
	scLingTongFashionWear.FashionId = &fashionId
	return scLingTongFashionWear
}

func BuildSCLingTongFashionUnload(fashionWear int32) *uipb.SCLingTongFashionUnload {
	lingTongFashionUnload := &uipb.SCLingTongFashionUnload{}
	lingTongFashionUnload.FashionId = &fashionWear
	return lingTongFashionUnload
}

func BuildSCLingTongFashionUpstar(fashionId int32, star int32, pro int32, sucess bool) *uipb.SCLingTongFashionUpstar {
	fahionUpstar := &uipb.SCLingTongFashionUpstar{}
	fahionUpstar.FashionId = &fashionId
	fahionUpstar.Level = &star
	fahionUpstar.LevelPro = &pro
	fahionUpstar.Result = &sucess
	return fahionUpstar
}

func BuildSCLingTongFashionActivate(obj *playerlingtong.PlayerLingTongFashionInfoObject) *uipb.SCLingTongFashionActive {
	scLingTongFashionActive := &uipb.SCLingTongFashionActive{}
	scLingTongFashionActive.FashionInfo = buildLingTongFashion(obj)
	return scLingTongFashionActive
}

func BuildSCLingTongFashionRemove(fashionId int32) *uipb.SCLingTongFashionRemove {
	lingTongFashionRemove := &uipb.SCLingTongFashionRemove{}
	lingTongFashionRemove.FashionId = &fashionId
	return lingTongFashionRemove

}

func BuildSCLingTongFashionTrialNotice(trialId int32, expireTime int64) *uipb.SCLingTongFashionTrialNotice {
	scLingTongFashionTrialNotice := &uipb.SCLingTongFashionTrialNotice{}
	scLingTongFashionTrialNotice.TrialFashionId = &trialId
	scLingTongFashionTrialNotice.ExpireTime = &expireTime
	return scLingTongFashionTrialNotice
}

func BuildSCLingTongFashionTrialOverdueNotice(trialId int32, overdueType int32) *uipb.SCLingTongFashionTrialOverdueNotice {
	scLingTongFashionTrialOverdueNotice := &uipb.SCLingTongFashionTrialOverdueNotice{}
	scLingTongFashionTrialOverdueNotice.TrialFashionId = &trialId
	scLingTongFashionTrialOverdueNotice.OverdueType = &overdueType
	return scLingTongFashionTrialOverdueNotice
}
