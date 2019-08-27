package player

import (
	advancedrewrewtemplate "fgame/fgame/game/welfare/advancedrew/rew/template"
	advancedrewrewtypes "fgame/fgame/game/welfare/advancedrew/rew/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

// 进阶奖励
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeRew, playerwelfare.ActivityObjInfoInitFunc(rewAdvancedInitInfo))
}

func rewAdvancedInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*advancedrewrewtypes.AdvancedRewInfo)
	info.RewRecord = []int32{}
	info.IsEmail = false
	info.PeriodChargeNum = 0

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*advancedrewrewtemplate.GroupTemplateRew)
	advancedType := groupTemp.GetAdvancedType()
	pl := obj.GetPlayer()
	info.AdvancedNum = welfare.GetSystemAdvancedNum(pl, advancedType)
	info.RewType = advancedType
}
