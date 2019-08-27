package player

import (
	advancedrewrewmaxtemplate "fgame/fgame/game/welfare/advancedrew/rew_max/template"
	advancedrewrewmaxtypes "fgame/fgame/game/welfare/advancedrew/rew_max/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

// 进阶奖励
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeRewMax, playerwelfare.ActivityObjInfoInitFunc(rewMaxAdvancedInitInfo))
}

func rewMaxAdvancedInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*advancedrewrewmaxtypes.AdvancedRewMaxInfo)
	info.RewRecord = []int32{}
	info.IsEmail = false
	info.PeriodChargeNum = 0

	pl := obj.GetPlayer()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*advancedrewrewmaxtemplate.GroupTemplateRewMax)

	advancedType := groupTemp.GetAdvancedType()
	info.AdvancedNum = welfare.GetSystemAdvancedNum(pl, advancedType)
	info.InitAdvancedNum = welfare.GetSystemAdvancedNum(pl, advancedType)
	info.RewType = advancedType
}
