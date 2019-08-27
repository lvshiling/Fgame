package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	cyclechargesinglemaxrewmultipletypes "fgame/fgame/game/welfare/cycle/charge_single_max_rew_multiple/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeMaxRewMultiple, reddot.HandlerFunc(handleRedDotCycleSingleMaxRewMultiple))
}

//每日单笔充值红点
func handleRedDotCycleSingleMaxRewMultiple(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	info := obj.GetActivityData().(*cyclechargesinglemaxrewmultipletypes.CycleSingleChargeMaxRewMultipleInfo)
	leftM := info.LeftCanReceiveRewards()
	if len(leftM) > 0 {
		isNotice = true
		return
	}

	return
}
