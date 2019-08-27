package handler

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	cyclechargesinglemaxrewmultipletypes "fgame/fgame/game/welfare/cycle/charge_single_max_rew_multiple/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeMaxRewMultiple, welfare.InfoGetHandlerFunc(cycleSingleMaxRewMultipleInfo))
}

//获取每日单笔充值请求逻辑
func cycleSingleMaxRewMultipleInfo(pl player.Player, groupId int32) (err error) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)

	var hadRewRecord map[int32]int32
	var canRewRecord map[int32]int32
	maxSingleNum := int32(0)
	if obj != nil {
		welfareManager.RefreshActivityDataByGroupId(groupId)

		info := obj.GetActivityData().(*cyclechargesinglemaxrewmultipletypes.CycleSingleChargeMaxRewMultipleInfo)
		canRewRecord = info.LeftCanReceiveRewards()
		hadRewRecord = info.NewReceiveRewRecord
		maxSingleNum = info.MaxSingleChargeNum
	}

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	scMsg := pbutil.BuildSCOpenActivityGetInfoCycleSingleMaxRewMultiple(groupId, startTime, endTime, hadRewRecord, maxSingleNum, canRewRecord)
	pl.SendMsg(scMsg)
	return
}
