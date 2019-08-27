package player

import (
	advancedrewexpendreturntemplate "fgame/fgame/game/welfare/advancedrew/expend_return/template"
	advancedrewexpendreturntypes "fgame/fgame/game/welfare/advancedrew/expend_return/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 进阶消耗返还
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeExpendReturn, playerwelfare.ActivityObjInfoInitFunc(expendReturnInitInfo))
}

func expendReturnInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	groupId := obj.GetGroupId()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*advancedrewexpendreturntemplate.GroupTemplateAdvancedExpendReturn)

	info := obj.GetActivityData().(*advancedrewexpendreturntypes.AdvancedExpendReturnInfo)
	info.DanNum = 0
	info.RewRecord = []int32{}
	info.RewType = groupTemp.GetAdvancedType()
	info.IsEmail = false
}
