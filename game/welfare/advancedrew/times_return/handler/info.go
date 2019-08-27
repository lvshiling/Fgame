package handler

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	advancedrewtimesreturntypes "fgame/fgame/game/welfare/advancedrew/times_return/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeTimesReturn, welfare.InfoGetHandlerFunc(timesReturnInfo))
}

//升阶次数返还信息
func timesReturnInfo(pl player.Player, groupId int32) (err error) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	err = welfareManager.RefreshActivityDataByGroupId(groupId)
	if err != nil {
		return
	}

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	obj := welfareManager.GetOpenActivity(groupId)
	var record []int32
	advancedTimes := int32(0)
	if obj != nil {
		info := obj.GetActivityData().(*advancedrewtimesreturntypes.AdvancedTimesReturnInfo)
		record = info.RewRecord
		advancedTimes = info.Times
	}

	scMsg := pbutil.BuildSCOpenActivityGetInfoAdvancedTimesReturn(groupId, startTime, endTime, record, advancedTimes)
	pl.SendMsg(scMsg)
	return
}
