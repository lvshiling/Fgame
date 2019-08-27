package handler

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	drewcostdrewtypes "fgame/fgame/game/welfare/drew/cost_drew/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeCostDrew, welfare.InfoGetHandlerFunc(getCostDrewInfo))
}

//获取消费抽奖请求逻辑
func getCostDrewInfo(pl player.Player, groupId int32) (err error) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	welfareManager.RefreshActivityDataByGroupId(groupId)

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	obj := welfareManager.GetOpenActivity(groupId)
	costNum := int32(0)
	leftTimes := int32(0)
	attendTimes := int32(0)
	rate := int32(0)
	if obj != nil {
		info := obj.GetActivityData().(*drewcostdrewtypes.LuckyCostDrewInfo)
		costNum = int32(info.GoldNum)
		leftTimes = info.LeftTimes
		attendTimes = info.AttendTimes
		rate = info.Ratio
	}

	logList := welfare.GetWelfareService().GetDrewLogByTime(groupId, 0)
	scMsg := pbutil.BuildSCOpenActivityGetInfoCostDrew(groupId, costNum, leftTimes, attendTimes, rate, logList, startTime, endTime)
	pl.SendMsg(scMsg)
	return
}
