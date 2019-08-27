package reddot

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/reddot/reddot"
	drewcommontypes "fgame/fgame/game/welfare/drew/common/types"
	drewtraytemplate "fgame/fgame/game/welfare/drew/tray/template"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeTray, reddot.HandlerFunc(handleRedDotTray))
}

//大转盘红点
func handleRedDotTray(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}

	groupTemp := groupInterface.(*drewtraytemplate.GroupTemplateDrewTray)
	needGold := groupTemp.GetLuckyDrewNeedGold(drewcommontypes.LuckyDrewTypeOnce)
	curGold := propertyManager.GetGold()
	if curGold < needGold {
		return
	}

	isNotice = true
	return
}
