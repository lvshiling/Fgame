package player

import (
	advancedrewtimesreturntemplate "fgame/fgame/game/welfare/advancedrew/times_return/template"
	advancedrewtimesreturntypes "fgame/fgame/game/welfare/advancedrew/times_return/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 进阶消耗返还
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeTimesReturn, playerwelfare.ActivityObjInfoInitFunc(timesReturnInitInfo))
}

func timesReturnInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	groupId := obj.GetGroupId()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*advancedrewtimesreturntemplate.GroupTemplateAdvancedTimesReturn)

	info := obj.GetActivityData().(*advancedrewtimesreturntypes.AdvancedTimesReturnInfo)
	info.Times = 0
	info.RewRecord = []int32{}
	info.RewType = groupTemp.GetAdvancedType()
	info.IsEmail = false
}
