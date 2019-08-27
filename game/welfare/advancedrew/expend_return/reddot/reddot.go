package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	advancedrewexpendreturntemplate "fgame/fgame/game/welfare/advancedrew/expend_return/template"
	advancedrewexpendreturntypes "fgame/fgame/game/welfare/advancedrew/expend_return/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeExpendReturn, reddot.HandlerFunc(handleRedDotAdvancedExpendReturn))
}

//升阶消耗返还红点
func handleRedDotAdvancedExpendReturn(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*advancedrewexpendreturntemplate.GroupTemplateAdvancedExpendReturn)
	info := obj.GetActivityData().(*advancedrewexpendreturntypes.AdvancedExpendReturnInfo)

	//对应的类型
	if info.RewType != groupTemp.GetAdvancedType() {
		return
	}

	//进阶返利
	for _, temp := range groupTemp.GetOpenTempMap() {
		if !info.IsCanReceiveRewards(temp.Value2) {
			continue
		}

		isNotice = true
	}

	return
}
