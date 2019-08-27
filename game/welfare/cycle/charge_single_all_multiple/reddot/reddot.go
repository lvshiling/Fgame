package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	chargesingleallmultipletemplate "fgame/fgame/game/welfare/cycle/charge_single_all_multiple/template"
	chargesingleallmultipletypes "fgame/fgame/game/welfare/cycle/charge_single_all_multiple/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeAllRew, reddot.HandlerFunc(handlerReddot))
}

func handlerReddot(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	info := obj.GetActivityData().(*chargesingleallmultipletypes.CycleChargeSingleAllMultipleInfo)
	groupId := obj.GetGroupId()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*chargesingleallmultipletemplate.GroupTemplateCycleSingleAllMultiple)
	canRewRecord := groupTemp.GetCanRewRecordMap(info.CycleDay, info.GetCanRewRecord())
	for _, value := range canRewRecord {
		if value > 0 {
			isNotice = true
			return
		}
	}
	isNotice = false
	return
}
