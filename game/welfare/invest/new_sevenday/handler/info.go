package handler

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	investnewsevendaytypes "fgame/fgame/game/welfare/invest/new_sevenday/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeInvest, welfaretypes.OpenActivityInvestSubTypeNewServenDay, welfare.InfoGetHandlerFunc(handleInvestNewSevenDayInfo))
}

func handleInvestNewSevenDayInfo(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeInvest
	subType := welfaretypes.OpenActivityInvestSubTypeNewServenDay

	// 校验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	info := obj.GetActivityData().(*investnewsevendaytypes.NewInvestDayInfo)

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	var record []int32

	scOpenActivityNewSevenDayInvestInfo := pbutil.BuildSCOpenActivityNewSevenDayInvestInfo(groupId, startTime, endTime, record, info.ReceiveMap, info.BuyTimeMap)
	pl.SendMsg(scOpenActivityNewSevenDayInvestInfo)

	scMsg := pbutil.BuildSCMergeActivitySingleChargeNotice(groupId, info.MaxSingleChargeNum)
	pl.SendMsg(scMsg)
	return
}
