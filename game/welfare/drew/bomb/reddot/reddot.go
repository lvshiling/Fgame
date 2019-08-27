package reddot

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	"fgame/fgame/game/reddot/reddot"
	drewbombtemplate "fgame/fgame/game/welfare/drew/bomb/template"
	drewcommontypes "fgame/fgame/game/welfare/drew/common/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeBombOre, reddot.HandlerFunc(handleRedDotBomb))
}

//炸矿红点
func handleRedDotBomb(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}

	groupTemp := groupInterface.(*drewbombtemplate.GroupTemplateDrewBomb)
	needGold := groupTemp.GetLuckyDrewNeedGold(drewcommontypes.LuckyDrewTypeOnce)
	curGold := propertyManager.GetGold()
	if curGold < needGold {
		return
	}

	isNotice = true
	return
}
