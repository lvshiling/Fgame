package player

import (
	feedbackcosttypes "fgame/fgame/game/welfare/feedback/cost/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 消费返利
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeCost, playerwelfare.ActivityObjInfoInitFunc(feedbackCostInitInfo))
}

func feedbackCostInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*feedbackcosttypes.FeedbackCostInfo)
	info.GoldNum = 0
	info.RewRecord = []int32{}
	info.IsEmail = false
}
