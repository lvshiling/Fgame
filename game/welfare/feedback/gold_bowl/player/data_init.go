package player

import (
	feedbackgoldbowltypes "fgame/fgame/game/welfare/feedback/gold_bowl/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 聚宝盆
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeGoldBowl, playerwelfare.ActivityObjInfoInitFunc(goldBowlInitInfo))
}

func goldBowlInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*feedbackgoldbowltypes.FeedbackGoldBowlInfo)
	info.GoldNum = 0
	info.IsEmail = false
}
