package player

import (
	feedbackchargereturnmultipletypes "fgame/fgame/game/welfare/feedback/charge_return_multiple/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 累充值返还
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeReturnMultiple, playerwelfare.ActivityObjInfoInitFunc(chargeReturnMultipleInitInfo))
}

func chargeReturnMultipleInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*feedbackchargereturnmultipletypes.FeedbackChargeReturnMultipleInfo)
	info.PeriodChargeNum = 0
	info.RewardCnt = 0
}
