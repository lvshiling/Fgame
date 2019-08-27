package player

import (
	feedbackchargesinglemaxrewtypes "fgame/fgame/game/welfare/feedback/charge_single_max_rew/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 单笔充值
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeSingleChagreMaxRew, playerwelfare.ActivityObjInfoInitFunc(singleMaxRewInitInfo))
}

func singleMaxRewInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*feedbackchargesinglemaxrewtypes.FeedbackSingleChargeMaxRewInfo)
	info.IsEmail = false
	info.MaxSingleChargeNum = 0
	info.ReceiveRewRecord = []int32{}
	info.CanRewRecord = []int32{}
}
