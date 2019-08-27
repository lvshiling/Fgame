package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	advancedrewrewtemplate "fgame/fgame/game/welfare/advancedrew/rew/template"
	advancedrewrewtypes "fgame/fgame/game/welfare/advancedrew/rew/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeRew, reddot.HandlerFunc(handleRedDotAdvancedRew))
}

//升阶祝福大放送红点
func handleRedDotAdvancedRew(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*advancedrewrewtemplate.GroupTemplateRew)
	info := obj.GetActivityData().(*advancedrewrewtypes.AdvancedRewInfo)
	if info.RewType != groupTemp.GetAdvancedType() {
		return
	}

	for _, temp := range groupTemp.GetOpenTempMap() {
		needAdvancedNum := temp.Value2
		needChargeNum := temp.Value3
		if !info.IsCanReceiveRewards(needAdvancedNum, needChargeNum) {
			continue
		}

		isNotice = true
		return
	}

	return
}
