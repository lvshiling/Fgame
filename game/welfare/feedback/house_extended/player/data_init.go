package player

import (
	feedbackhouseextendedtypes "fgame/fgame/game/welfare/feedback/house_extended/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 房产活动
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeHouseExtended, playerwelfare.ActivityObjInfoInitFunc(feedbackHouseExtendedInitInfo))
}

func feedbackHouseExtendedInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*feedbackhouseextendedtypes.FeedbackHouseExtendedInfo)
	info.ActivateChargeNum = 0
	info.IsActivateGift = false
	info.UplevelChargeNum = 0
	info.CurUplevelGiftLevel = 0
	info.IsUplevelGift = false
}
