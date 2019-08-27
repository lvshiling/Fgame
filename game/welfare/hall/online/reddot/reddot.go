package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	hallonlinetemplate "fgame/fgame/game/welfare/hall/online/template"
	hallonlinetypes "fgame/fgame/game/welfare/hall/online/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeOnline, reddot.HandlerFunc(handleRedDotWelfareOnline))
}

//福利大厅在线红点
func handleRedDotWelfareOnline(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()

	//在线奖励
	onlineInfo := obj.GetActivityData().(*hallonlinetypes.WelfareOnlineInfo)
	onlineTime := pl.GetTodayOnlineTime()
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}

	groupTemp := groupInterface.(*hallonlinetemplate.GroupTemplateWelfareOnline)
	curMaxTimes := groupTemp.GetOpenActivityWelfareOnlineDrewTimes(onlineTime)
	if onlineInfo.DrawTimes >= curMaxTimes {
		return
	}

	isNotice = true
	return
}
