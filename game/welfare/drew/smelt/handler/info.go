package handler

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	smelttemplate "fgame/fgame/game/welfare/drew/smelt/template"
	smelttypes "fgame/fgame/game/welfare/drew/smelt/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeSmelt, welfare.InfoGetHandlerFunc(smeltInfoHandle))
}

func smeltInfoHandle(pl player.Player, groupId int32) (err error) {
	typ := welfaretypes.OpenActivityTypeMergeDrew
	subType := welfaretypes.OpenActivityDrewSubTypeSmelt
	//检验活动
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*smelttemplate.GroupTemplateSmelt)

	var hadRewRecord []int32
	var useItemNum int32
	var canRewRecord int32
	if obj != nil {
		itemNeedNum := groupTemp.GetNeedItemNum()
		info := obj.GetActivityData().(*smelttypes.SmeltInfo)
		hadRewRecord = []int32{info.HasReceiveNum}
		canRewRecord = info.GetRemainCanReceiveRecord(itemNeedNum)
		useItemNum = info.Num
	}
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	scMsg := pbutil.BuildSCOpenActivityGetInfoSmelt(groupId, startTime, endTime, hadRewRecord, canRewRecord, useItemNum)
	pl.SendMsg(scMsg)
	return
}
