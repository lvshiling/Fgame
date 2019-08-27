package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	drewcostdrewtypes "fgame/fgame/game/welfare/drew/cost_drew/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeCostDrew, reddot.HandlerFunc(handleRedDotCostDrew))
}

//消费抽奖
func handleRedDotCostDrew(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	info := obj.GetActivityData().(*drewcostdrewtypes.LuckyCostDrewInfo)
	if info.LeftTimes > 0 {
		isNotice = true
		return
	}

	// groupInteface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	// if groupInteface == nil {
	// 	return
	// }
	// groupTemp := groupInteface.(*drewcostdrewtemplate.GroupTemplateCostDrew)
	// needGold := groupTemp.GetCostDrewNeedGold()
	// propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	// if !propertyManager.HasEnoughGold(needGold, false) {
	// 	return
	// }

	return
}
