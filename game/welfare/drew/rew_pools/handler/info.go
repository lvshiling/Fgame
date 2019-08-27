package handler

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	rewpoolstypes "fgame/fgame/game/welfare/drew/rew_pools/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeRewPools, welfare.InfoGetHandlerFunc(rewPoolsInfoHandle))
}

func rewPoolsInfoHandle(pl player.Player, groupId int32) (err error) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)

	var position int32
	var backTimes int32
	if obj != nil {
		// welfareManager.RefreshActivityDataByGroupId(groupId)
		info := obj.GetActivityData().(*rewpoolstypes.RewPoolsInfo)
		position = info.Position
		backTimes = info.BackTimes
	}
	logList := welfare.GetWelfareService().GetDrewLogByTime(groupId, 0)
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	scMsg := pbutil.BuildSCOpenActivityGetInfoRewPools(groupId, startTime, endTime, position, backTimes, logList)
	pl.SendMsg(scMsg)
	return
}
