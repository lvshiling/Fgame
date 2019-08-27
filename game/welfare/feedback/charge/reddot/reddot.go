package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	feedbackchargetemplate "fgame/fgame/game/welfare/feedback/charge/template"
	feedbackchargetypes "fgame/fgame/game/welfare/feedback/charge/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeCharge, reddot.HandlerFunc(handleRedDotOpenFeedback))
}

//开服累充红点
func handleRedDotOpenFeedback(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()

	info := obj.GetActivityData().(*feedbackchargetypes.FeedbackChargeInfo)
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}

	//可领取累充
	groupTemp := groupInterface.(*feedbackchargetemplate.GroupTemplateCharge)
	for _, temp := range groupTemp.GetOpenTempMap() {
		needCharge := temp.Value1
		if !info.IsCanReceiveRewards(needCharge) {
			continue
		}

		isNotice = true
		return
	}

	return
}
