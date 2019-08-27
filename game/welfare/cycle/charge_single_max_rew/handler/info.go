package handler

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	cyclechargesinglemaxrewtypes "fgame/fgame/game/welfare/cycle/charge_single_max_rew/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeMaxRew, welfare.InfoGetHandlerFunc(cycleSingleMaxRewInfo))
}

//获取每日单笔充值请求逻辑
func cycleSingleMaxRewInfo(pl player.Player, groupId int32) (err error) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)

	var hadRewRecord []int32
	var canRewRecord []int32
	maxSingleNum := int32(0)
	if obj != nil {
		welfareManager.RefreshActivityDataByGroupId(groupId)

		info := obj.GetActivityData().(*cyclechargesinglemaxrewtypes.CycleSingleChargeMaxRewInfo)
		canRewRecord = info.CanRewRecord
		hadRewRecord = info.ReceiveRewRecord
		maxSingleNum = info.MaxSingleChargeNum
	}

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	scMsg := pbutil.BuildSCOpenActivityGetInfoCycleSingleMaxRew(groupId, startTime, endTime, hadRewRecord, maxSingleNum, canRewRecord)
	pl.SendMsg(scMsg)
	return
}
