package handler

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	feedbackchargereturntypes "fgame/fgame/game/welfare/feedback/charge_return/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeReturn, welfare.InfoGetHandlerFunc(handlerChargeReturnInfo))
}

//充值返还
func handlerChargeReturnInfo(pl player.Player, groupId int32) (err error) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	obj := welfareManager.GetOpenActivity(groupId)
	var record []int32
	isReturn := false
	if obj != nil {
		info := obj.GetActivityData().(*feedbackchargereturntypes.FeedbackChargeReturnInfo)
		isReturn = info.IsReturn
	}

	scMsg := pbutil.BuildSCOpenActivityGetInfoChargeReturn(groupId, startTime, endTime, record, isReturn)
	pl.SendMsg(scMsg)
	return
}
