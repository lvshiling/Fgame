package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	investsevendaytemplate "fgame/fgame/game/welfare/invest/sevenday/template"
	investsevendaytypes "fgame/fgame/game/welfare/invest/sevenday/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeInvest, welfaretypes.OpenActivityInvestSubTypeServenDay, reddot.HandlerFunc(handleRedDotInvestDay))
}

//七日投资红点
func handleRedDotInvestDay(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	info := obj.GetActivityData().(*investsevendaytypes.InvestDayInfo)
	groupId := obj.GetGroupId()
	// groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	// if groupInterface == nil {
	// 	return
	// }
	// groupTemp := groupInterface.(*investsevendaytemplate.GroupTemplateInvestDay)
	// // 未购买-元宝足够
	// if info.BuyTime <= 0 {
	// 	needGold := int64(groupTemp.GetInvestDayNeedGold())
	// 	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	// 	if !propertyManager.HasEnoughGold(needGold, false) {
	// 		return
	// 	}

	// 	isNotice = true
	// 	return
	// }

	// 七天是否已经全部领取完
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*investsevendaytemplate.GroupTemplateInvestDay)
	maxRewardDay := groupTemp.GetInvestDayMaxRewardsLevel()

	// 已购买可领取
	if !info.IsCanReceive(maxRewardDay) {
		return
	}

	isNotice = true
	return
}
