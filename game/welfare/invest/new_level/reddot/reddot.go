package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	investnewleveltemplate "fgame/fgame/game/welfare/invest/new_level/template"
	investnewleveltypes "fgame/fgame/game/welfare/invest/new_level/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeInvest, welfaretypes.OpenActivityInvestSubTypeNewLevel, reddot.HandlerFunc(handleRedDotInvestNewLevel))
}

//投资计划红点
func handleRedDotInvestNewLevel(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*investnewleveltemplate.GroupTemplateInvestNewLevel)
	curCondition := pl.GetLevel()
	info := obj.GetActivityData().(*investnewleveltypes.InvestNewLevelInfo)

	// //有钱可以买
	// if len(info.InvestBuyInfoMap) == 0 {
	// propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	// 	for typeTemp := investnewleveltypes.MinType; typeTemp <= investnewleveltypes.MaxType; typeTemp++ {
	// 		needGold := int64(groupTemp.GetInvestLevelNeedGold(typeTemp))
	// 		if propertyManager.HasEnoughGold(needGold, false) {
	// 			isNotice = true
	// 			return
	// 		}
	// 	}
	// 	return false
	// }

	for investType, recordList := range info.InvestBuyInfoMap {
		tempM := groupTemp.GetInvestLevelTempMByType(investType)
		if len(recordList) >= len(tempM) {
			return
		}
		for rewLev := range tempM {
			if rewLev > curCondition {
				continue
			}
			if !info.IsCanReceiveRew(investType, rewLev) {
				continue
			}
			isNotice = true
			return
		}
	}

	return
}
