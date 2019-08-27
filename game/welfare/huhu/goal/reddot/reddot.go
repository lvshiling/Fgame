package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	goaltemplate "fgame/fgame/game/welfare/huhu/goal/template"
	goaltypes "fgame/fgame/game/welfare/huhu/goal/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeHuHu, welfaretypes.OpenActivitySpecialSubTypeGoal, reddot.HandlerFunc(handleRedDotGoal))
}

//目标红点
func handleRedDotGoal(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*goaltemplate.GroupTemplateGoal)

	info := obj.GetActivityData().(*goaltypes.GoalInfo)
	rewTempList := groupTemp.GetCanRewTemplate(info.GoalCount, info.RewRecordMap)
	if len(rewTempList) <= 0 {
		return
	}

	isNotice = true
	return
}
