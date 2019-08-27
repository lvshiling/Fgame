package player

import (
	feedbackchargereturnleveltypes "fgame/fgame/game/welfare/feedback/charge_return_level/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 充值返还
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeReturnLevel, playerwelfare.ActivityObjInfoInitFunc(chargeReturnLevelInitInfo))
}

func chargeReturnLevelInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*feedbackchargereturnleveltypes.FeedbackChargeReturnLevelInfo)
	info.IsReturn = false
}
