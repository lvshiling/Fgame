package handler

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	feedbackchargereturnmultipletypes "fgame/fgame/game/welfare/feedback/charge_return_multiple/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeReturnMultiple, welfare.InfoGetHandlerFunc(handlerChargeReturnMultipleInfo))
}

//充值返还
func handlerChargeReturnMultipleInfo(pl player.Player, groupId int32) (err error) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	obj := welfareManager.GetOpenActivity(groupId)
	periodChargeNum := int32(0)
	rewardCnt := int32(0)
	var record []int32
	if obj != nil {
		err = welfareManager.RefreshActivityDataByGroupId(groupId)
		if err != nil {
			return
		}
		info := obj.GetActivityData().(*feedbackchargereturnmultipletypes.FeedbackChargeReturnMultipleInfo)
		periodChargeNum = info.PeriodChargeNum
		rewardCnt = info.RewardCnt
	}

	scMsg := pbutil.BuildSCOpenActivityGetInfoChargeReturnMultiple(groupId, startTime, endTime, record, periodChargeNum, rewardCnt)
	pl.SendMsg(scMsg)
	return
}
