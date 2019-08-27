package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	cyclechargetemplate "fgame/fgame/game/welfare/cycle/charge/template"
	cyclechargetypes "fgame/fgame/game/welfare/cycle/charge/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeCharge, reddot.HandlerFunc(handleRedDotCycleCharge))
}

//每日充值红点
func handleRedDotCycleCharge(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()

	info := obj.GetActivityData().(*cyclechargetypes.CycleChargeInfo)
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*cyclechargetemplate.GroupTemplateCycleCharge)
	tempList := groupTemp.GetCurDayTempList(info.CycleDay)
	for _, temp := range tempList {
		needCharge := temp.Value2
		if !info.IsCanReceiveRewards(needCharge) {
			continue
		}

		isNotice = true
		return
	}

	return
}
