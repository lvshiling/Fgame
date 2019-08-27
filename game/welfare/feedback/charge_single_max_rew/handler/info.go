package handler

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	feedbackchargesinglemaxrewtypes "fgame/fgame/game/welfare/feedback/charge_single_max_rew/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeSingleChagreMaxRew, welfare.InfoGetHandlerFunc(singleChargeMaxRewInfo))
}

//获取单笔充值信息（最近档次）
func singleChargeMaxRewInfo(pl player.Player, groupId int32) (err error) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)

	var hadRewRecord []int32
	var canRewRecord []int32
	maxSingleNum := int32(0)

	if obj != nil {
		welfareManager.RefreshActivityDataByGroupId(groupId)

		info := obj.GetActivityData().(*feedbackchargesinglemaxrewtypes.FeedbackSingleChargeMaxRewInfo)
		canRewRecord = info.CanRewRecord
		hadRewRecord = info.ReceiveRewRecord
		maxSingleNum = info.MaxSingleChargeNum
	}

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	scMsg := pbutil.BuildSCOpenActivityGetInfoFeedbackSingleMaxRew(groupId, startTime, endTime, hadRewRecord, maxSingleNum, canRewRecord)
	pl.SendMsg(scMsg)
	return
}
