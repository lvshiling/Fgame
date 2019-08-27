package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	feedbackhouseinvesttemplate "fgame/fgame/game/welfare/feedback/house_invest/template"
	feedbackhouseinvesttypes "fgame/fgame/game/welfare/feedback/house_invest/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeHouseInvest, reddot.HandlerFunc(handleRedDotOpenFeedbackHouseInvest))
}

//开服-房产投资红点
func handleRedDotOpenFeedbackHouseInvest(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()

	info := obj.GetActivityData().(*feedbackhouseinvesttypes.FeedbackHouseInvestInfo)
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*feedbackhouseinvesttemplate.GroupTemplateHouseInvest)

	if info.IsSell {
		return
	}

	if !info.IsActivity {
		firstOpenTemp := groupTemp.GetOpenActivityHouseInvest(0)
		if firstOpenTemp == nil {
			return
		}
		if info.ChargeNum >= groupTemp.GetOpenActivityHouseInvestChargeNum(0) {
			isNotice = true
		}
		return
	}

	openTemp := groupTemp.GetOpenActivityHouseInvest(info.DecorDays + 1)
	if openTemp == nil {
		isNotice = true
		return
	}

	if !info.IsCurDayDecor && info.CurDayChargeNum >= groupTemp.GetOpenActivityHouseInvestChargeNum(info.DecorDays+1) {
		isNotice = true
		return
	}

	return
}
