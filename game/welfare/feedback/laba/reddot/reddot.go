package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	feedbacklabatypes "fgame/fgame/game/welfare/feedback/laba/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeGoldLaBa, reddot.HandlerFunc(handleRedDotOpenFeedbackLaba))
}

//拉霸红点
func handleRedDotOpenFeedbackLaba(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*feedbacklabatypes.FeedbackGoldLaBaInfo)

	nextTimes := info.Times + 1
	nextLaBaTemp := welfaretemplate.GetWelfareTemplateService().GetGoldLabaTemplate(groupId, nextTimes)
	if nextLaBaTemp == nil {
		return
	}
	if nextLaBaTemp.InvestmentRecharge > info.ChargeNum {
		return
	}

	isNotice = true
	return
}
