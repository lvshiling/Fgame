package player

import (
	feedbackchargecycletypes "fgame/fgame/game/welfare/feedback/charge_cycle/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 充值返利
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeCycleCharge, playerwelfare.ActivityObjInfoInitFunc(feedbackCycleChargeInitInfo))
}

func feedbackCycleChargeInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*feedbackchargecycletypes.FeedbackCycleChargeInfo)
	info.CycleDay = welfarelogic.CountCurActivityDay(groupId)
	info.DayNum = 0
	info.CurDayChargeNum = 0
	info.IsReceiveDayRew = false
	info.RewRecord = []int32{}
}
