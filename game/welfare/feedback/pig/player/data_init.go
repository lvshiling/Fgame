package player

import (
	feedbackpigtypes "fgame/fgame/game/welfare/feedback/pig/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 养金猪
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeGoldPig, playerwelfare.ActivityObjInfoInitFunc(goldPigInitInfo))
}

func goldPigInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*feedbackpigtypes.FeedbackGoldPigInfo)
	info.ChargeGold = 0
	info.CostGold = 0
	info.CurCondition = 0
	info.IsEmail = false
}
