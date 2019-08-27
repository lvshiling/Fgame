package reddot

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	drewsmasheggtemplate "fgame/fgame/game/welfare/drew/smash_egg/template"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
)

func init() {
	// reddot.Register(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeSmashEgg, reddot.HandlerFunc(handleRedDotSmashEgg))
}

//砸金蛋红点
func handleRedDotSmashEgg(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*drewsmasheggtemplate.GroupTemplateSmashEgg)
	needGold := groupTemp.GetNeedGold()
	needBindGold := groupTemp.GetNeedBindGold()
	needSilver := groupTemp.GetNeedSilver()
	flag := propertyManager.HasEnoughCost(int64(needBindGold), int64(needGold), int64(needSilver))
	if !flag {
		return
	}

	isNotice = true
	return
}
