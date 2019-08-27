package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	feedbackcosttypes "fgame/fgame/game/welfare/feedback/cost/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeCost, reddot.HandlerFunc(handleRedDotOpenFeedbackCost))
}

//开服消费红点
func handleRedDotOpenFeedbackCost(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()

	//累充
	info := obj.GetActivityData().(*feedbackcosttypes.FeedbackCostInfo)
	tempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
	for _, temp := range tempList {
		needCost := temp.Value1
		if !info.IsCanReceiveRewards(needCost) {
			continue
		}
		isNotice = true
		return
	}

	return
}
