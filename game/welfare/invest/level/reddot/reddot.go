package reddot

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	investleveltemplate "fgame/fgame/game/welfare/invest/level/template"
	investleveltypes "fgame/fgame/game/welfare/invest/level/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
)

func init() {
	// reddot.Register(welfaretypes.OpenActivityTypeInvest, welfaretypes.OpenActivityInvestSubTypeLevel, reddot.HandlerFunc(handleRedDotInvest))
}

//投资计划红点
func handleRedDotInvest(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*investleveltemplate.GroupTemplateInvestLevel)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	curCondition := pl.GetLevel()
	info := obj.GetActivityData().(*investleveltypes.InvestLevelInfo)

	// 初级投资
	juniorType := investleveltypes.InvesetLevelTypeJunior
	juniorRecord, ok := info.InvestBuyInfoMap[juniorType]
	if !ok {
		//初级元宝条件
		needGold := int64(groupTemp.GetInvestLevelNeedGold(juniorType))
		if propertyManager.HasEnoughGold(needGold, false) {
			isNotice = true
			return
		}
	} else {
		rewTempList := groupTemp.GetInvestLevelTempList(juniorType, juniorRecord, curCondition)
		if len(rewTempList) > 0 {
			isNotice = true
			return
		}
	}

	//高级投资
	seniorType := investleveltypes.InvesetLevelTypeSenior
	seniorRecord, ok := info.InvestBuyInfoMap[seniorType]
	if !ok {
		//高级元宝条件
		needGold := int64(groupTemp.GetInvestLevelNeedGold(seniorType))
		if propertyManager.HasEnoughGold(needGold, false) {
			isNotice = true
			return
		}
	} else {
		rewTempList := groupTemp.GetInvestLevelTempList(seniorType, seniorRecord, curCondition)
		if len(rewTempList) > 0 {
			isNotice = true
			return
		}
	}
	return
}
