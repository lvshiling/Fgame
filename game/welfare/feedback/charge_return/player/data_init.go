package player

import (
	feedbackchargereturntypes "fgame/fgame/game/welfare/feedback/charge_return/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 充值返还
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeReturn, playerwelfare.ActivityObjInfoInitFunc(chargeReturnInitInfo))
}

func chargeReturnInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*feedbackchargereturntypes.FeedbackChargeReturnInfo)
	info.IsReturn = false
}
