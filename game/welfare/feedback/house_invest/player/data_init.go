package player

import (
	feedbackhouseinvesttypes "fgame/fgame/game/welfare/feedback/house_invest/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 房产投资
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeHouseInvest, playerwelfare.ActivityObjInfoInitFunc(feedbackHouseInvestInitInfo))
}

func feedbackHouseInvestInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*feedbackhouseinvesttypes.FeedbackHouseInvestInfo)
	info.ChargeNum = 0
	info.IsActivity = false
	info.CurDayChargeNum = 0
	info.IsCurDayDecor = false
	info.DecorDays = 0
	info.IsSell = false
}
