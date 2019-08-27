package info

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	feedbackchargedoubletypes "fgame/fgame/game/welfare/feedback/charge_double/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeDouble, welfare.InfoGetHandlerFunc(handlerChargeDoubleInfo))
}

//充值翻倍
func handlerChargeDoubleInfo(pl player.Player, groupId int32) (err error) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	obj := welfareManager.GetOpenActivity(groupId)
	var record []int32
	if obj != nil {
		info := obj.GetActivityData().(*feedbackchargedoubletypes.FeedbackChargeDoubleInfo)
		record = info.Record
	}

	scMsg := pbutil.BuildSCOpenActivityGetInfo(groupId, startTime, endTime, record)
	pl.SendMsg(scMsg)
	return
}
