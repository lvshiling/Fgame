package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	feedbackchargereturnmultipletemplate "fgame/fgame/game/welfare/feedback/charge_return_multiple/template"
	feedbackchargereturnmultipletypes "fgame/fgame/game/welfare/feedback/charge_return_multiple/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeReturnMultiple, reddot.HandlerFunc(handleRedDotChargeReturnMultiple))
}

//循环充值红点
func handleRedDotChargeReturnMultiple(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*feedbackchargereturnmultipletypes.FeedbackChargeReturnMultipleInfo)
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}

	groupTemp := groupInterface.(*feedbackchargereturnmultipletemplate.GroupTemplateChargeReturnMultiple)
	totalCnt := info.PeriodChargeNum / groupTemp.GetPerChargeNum()
	rewardLimitCnt := groupTemp.GetRewardLimitCnt()
	if rewardLimitCnt <= info.RewardCnt {
		return
	}

	if totalCnt <= info.RewardCnt {
		return
	}

	isNotice = true
	return
}
