package handler

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	developfamoustypes "fgame/fgame/game/welfare/develop/famous/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeDevelop, welfaretypes.OpenActivityDefaultSubTypeDefault, welfare.InfoGetHandlerFunc(handlerFameInfo))
}

//名人普信息请求
func handlerFameInfo(pl player.Player, groupId int32) (err error) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	obj := welfareManager.GetOpenActivity(groupId)
	var record []int32
	feedTimes := make(map[int32]int32)
	favorableNum := int32(0)
	dayFavorableNum := int32(0)
	if obj != nil {
		info := obj.GetActivityData().(*developfamoustypes.DevelopFameInfo)
		record = info.RewRecord
		favorableNum = info.FavorableNum
		feedTimes = info.FeedTimesMap
		dayFavorableNum = info.DayFavorableNum
	}

	scMsg := pbutil.BuildSCOpenActivityGetInfoDevelopFame(groupId, startTime, endTime, record, favorableNum, dayFavorableNum, feedTimes)
	pl.SendMsg(scMsg)
	return
}
