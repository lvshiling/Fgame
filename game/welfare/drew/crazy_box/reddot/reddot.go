package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	drewcrazyboxtemplate "fgame/fgame/game/welfare/drew/crazy_box/template"
	drewcrazyboxtypes "fgame/fgame/game/welfare/drew/crazy_box/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeCrazyBox, reddot.HandlerFunc(handleRedDotCrazyBox))
}

//疯狂宝箱
func handleRedDotCrazyBox(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()

	info := obj.GetActivityData().(*drewcrazyboxtypes.CrazyBoxInfo)
	if info.GoldNum <= 0 {
		return
	}

	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*drewcrazyboxtemplate.GroupTemplateCrazyBox)

	_, boxLeftTimes := groupTemp.GetCrazyBoxArg(info.AttendTimes)
	totalTimes := groupTemp.GetCrazyBoxTotalTimes(info.GoldNum)
	leftTimes := totalTimes - info.AttendTimes

	if boxLeftTimes <= 0 {
		return
	}

	if leftTimes <= 0 {
		return
	}

	isNotice = true
	return
}
