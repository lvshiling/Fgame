package player

import (
	feedbacklabatypes "fgame/fgame/game/welfare/feedback/laba/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 返利-元宝拉霸
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeGoldLaBa, playerwelfare.ActivityObjInfoInitFunc(feedbackGoldLaBaInitInfo))
}

func feedbackGoldLaBaInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*feedbacklabatypes.FeedbackGoldLaBaInfo)
	info.Times = 0
	info.ChargeNum = 0
}
