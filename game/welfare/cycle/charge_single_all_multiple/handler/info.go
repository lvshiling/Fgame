package handler

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	chargesingleallmultipletemplate "fgame/fgame/game/welfare/cycle/charge_single_all_multiple/template"
	chargesingleallmultipletypes "fgame/fgame/game/welfare/cycle/charge_single_all_multiple/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeAllRew, welfare.InfoGetHandlerFunc(handlerInfo))
}

func handlerInfo(pl player.Player, groupId int32) (err error) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	obj := welfareManager.GetOpenActivity(groupId)
	if obj == nil {
		return
	}

	welfareManager.RefreshActivityDataByGroupId(groupId)
	info := obj.GetActivityData().(*chargesingleallmultipletypes.CycleChargeSingleAllMultipleInfo)

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*chargesingleallmultipletemplate.GroupTemplateCycleSingleAllMultiple)

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	canRewRecord := groupTemp.GetCanRewRecordMap(info.CycleDay, info.GetCanRewRecord())
	scMsg := pbutil.BuildSCOpenActivityGetInfoCycleSingleAllRewMultiple(groupId, startTime, endTime, info.SingleChargeRecord, canRewRecord)
	pl.SendMsg(scMsg)
	return
}
