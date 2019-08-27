package player

import (
	feedbackchargedoubletypes "fgame/fgame/game/welfare/feedback/charge_double/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 充值翻倍
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeDouble, playerwelfare.ActivityObjInfoInitFunc(returnDoubelInitInfo))
}

func returnDoubelInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*feedbackchargedoubletypes.FeedbackChargeDoubleInfo)
	info.Record = []int32{}
}
