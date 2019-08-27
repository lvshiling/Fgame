package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	cyclechargesinglemaxrewtypes "fgame/fgame/game/welfare/cycle/charge_single_max_rew/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeMaxRew, reddot.HandlerFunc(handleRedDotCycleSingleMaxRew))
}

//每日单笔充值红点
func handleRedDotCycleSingleMaxRew(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	info := obj.GetActivityData().(*cyclechargesinglemaxrewtypes.CycleSingleChargeMaxRewInfo)
	if len(info.CanRewRecord) > 0 {
		isNotice = true
		return
	}

	return
}
