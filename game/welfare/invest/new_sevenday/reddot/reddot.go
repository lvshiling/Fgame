package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	investnewsevendaytemplate "fgame/fgame/game/welfare/invest/new_sevenday/template"
	investnewsevendaytypes "fgame/fgame/game/welfare/invest/new_sevenday/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeInvest, welfaretypes.OpenActivityInvestSubTypeNewServenDay, reddot.HandlerFunc(handleRedDotInvestDay))
}

//七日投资红点
func handleRedDotInvestDay(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	info := obj.GetActivityData().(*investnewsevendaytypes.NewInvestDayInfo)
	groupId := obj.GetGroupId()

	// 七天是否已经全部领取完
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*investnewsevendaytemplate.GroupTemplateNewInvestDay)
	maxRewardDay := groupTemp.GetInvestDayMaxDay()

	// 已购买可领取
	if !info.IsCanReceiveRettot(maxRewardDay) {
		return
	}

	isNotice = true
	return
}
