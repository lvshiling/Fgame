package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	feedbackchargesingletypes "fgame/fgame/game/welfare/feedback/charge_single/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeSingleChagre, reddot.HandlerFunc(handleRedDotSingleCharge))
}

//合服单笔充值红点
func handleRedDotSingleCharge(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()

	//单笔充值
	info := obj.GetActivityData().(*feedbackchargesingletypes.FeedbackSingleChargeInfo)
	tempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTemplateByGroup(groupId)
	for _, temp := range tempList {
		needNum := temp.Value1
		if !info.IsCanReceiveRewards(needNum) {
			continue
		}

		isNotice = true
		return
	}

	return
}
