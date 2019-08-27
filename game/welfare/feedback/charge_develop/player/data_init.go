package player

import (
	feedbackchargedeveloptypes "fgame/fgame/game/welfare/feedback/charge_develop/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 返利-培养
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeDevelop, playerwelfare.ActivityObjInfoInitFunc(developInitInfo))
}

func developInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*feedbackchargedeveloptypes.FeedbackDevelopInfo)
	info.ActivateChargeNum = 0
	info.IsDead = false
	info.IsFeed = false
	info.IsReceiveRew = false
	info.IsActivate = false
	info.TodayCostNum = 0
	info.FeedTimes = 0
	info.IsEndMail = false
}
