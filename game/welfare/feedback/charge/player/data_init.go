package player

import (
	feedbackchargetypes "fgame/fgame/game/welfare/feedback/charge/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 充值返利
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeCharge, playerwelfare.ActivityObjInfoInitFunc(feedbackChargeInitInfo))
}

func feedbackChargeInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*feedbackchargetypes.FeedbackChargeInfo)
	info.GoldNum = 0
	info.RewRecord = []int32{}
	info.IsEmail = false
}
