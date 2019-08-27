package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	smelttemplate "fgame/fgame/game/welfare/drew/smelt/template"
	smelttypes "fgame/fgame/game/welfare/drew/smelt/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeSmelt, reddot.HandlerFunc(handleRedDotSmelt))
}

func handleRedDotSmelt(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()
	welfareTemplateService := welfaretemplate.GetWelfareTemplateService()
	groupInterface := welfareTemplateService.GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}

	info := obj.GetActivityData().(*smelttypes.SmeltInfo)
	groupTemp := groupInterface.(*smelttemplate.GroupTemplateSmelt)
	needNum := groupTemp.GetNeedItemNum()
	remainNum := info.GetRemainCanReceiveRecord(needNum)
	if remainNum > 0 {
		return true
	}

	return false
}
