package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	cyclechargesingletemplate "fgame/fgame/game/welfare/cycle/charge_single/template"
	cyclechargesingletypes "fgame/fgame/game/welfare/cycle/charge_single/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleCharge, reddot.HandlerFunc(handleRedDotCycleSingle))
}

//每日单笔充值红点
func handleRedDotCycleSingle(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*cyclechargesingletemplate.GroupTemplateCycleSingle)
	curCycDay := welfarelogic.CountCycleDay(groupId)
	activityTemplateList := groupTemp.GetCurDayTempList(curCycDay)
	if len(activityTemplateList) == 0 {
		return
	}
	info := obj.GetActivityData().(*cyclechargesingletypes.CycleSingleChargeInfo)
	for _, activityTemplate := range activityTemplateList {
		if info.IsCanReceiveRewards(activityTemplate.Value2) {
			return true
		}
	}

	return false
}
