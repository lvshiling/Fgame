package player

import (
	feedbackchargesingletypes "fgame/fgame/game/welfare/feedback/charge_single/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 单笔充值
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeSingleChagre, playerwelfare.ActivityObjInfoInitFunc(singleInitInfo))
}

func singleInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*feedbackchargesingletypes.FeedbackSingleChargeInfo)
	info.IsEmail = false
	info.MaxSingleChargeNum = 0
	info.RewRecord = []int32{}
}
