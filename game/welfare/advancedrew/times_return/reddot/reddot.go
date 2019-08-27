package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	advancedrewtimesreturntemplate "fgame/fgame/game/welfare/advancedrew/times_return/template"
	advancedrewtimesreturntypes "fgame/fgame/game/welfare/advancedrew/times_return/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeAdvancedRew, welfaretypes.OpenActivityAdvancedRewSubTypeTimesReturn, reddot.HandlerFunc(handleRedDotAdvancedTimesReturn))
}

//升阶次数返还红点
func handleRedDotAdvancedTimesReturn(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*advancedrewtimesreturntemplate.GroupTemplateAdvancedTimesReturn)
	info := obj.GetActivityData().(*advancedrewtimesreturntypes.AdvancedTimesReturnInfo)

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
