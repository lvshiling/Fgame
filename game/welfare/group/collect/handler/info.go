package handler

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	groupcollecttypes "fgame/fgame/game/welfare/group/collect/types"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeGroup, welfaretypes.OpenActivityGroupSubTypeCollectPoker, welfare.InfoGetHandlerFunc(getCollectPokerInfo))
}

//获取卡牌收集信息
func getCollectPokerInfo(pl player.Player, groupId int32) (err error) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	obj := welfareManager.GetOpenActivity(groupId)
	var hadCollectRecord []int32
	if obj != nil {
		info := obj.GetActivityData().(*groupcollecttypes.CollectRewInfo)
		hadCollectRecord = info.HadPokerList
	}

	scMsg := pbutil.BuildSCOpenActivityGetInfo(groupId, startTime, endTime, hadCollectRecord)
	pl.SendMsg(scMsg)
	return
}
