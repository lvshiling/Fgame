package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	feedbackchargecycletemplate "fgame/fgame/game/welfare/feedback/charge_cycle/template"
	feedbackchargecycletypes "fgame/fgame/game/welfare/feedback/charge_cycle/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeCycleCharge, reddot.HandlerFunc(handleRedDotOpenFeedbackCycle))
}

//开服连续充值红点
func handleRedDotOpenFeedbackCycle(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()

	info := obj.GetActivityData().(*feedbackchargecycletypes.FeedbackCycleChargeInfo)
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*feedbackchargecycletemplate.GroupTemplateCycleCharge)

	//当天充值
	if !info.IsReceiveDayRew {
		needCharge := groupTemp.GetDayRewCondition(info.CycleDay)
		if info.CurDayChargeNum >= needCharge {
			isNotice = true
			return
		}
	}

	// 累计充值
	for needDay, _ := range groupTemp.GetEndRewTempMap() {
		if !info.IsCanReceiveCountDay(needDay) {
			continue
		}

		isNotice = true
		return
	}

	return
}
