package player

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	chargelimitlogic "fgame/fgame/game/welfare/rewards/charge_limit/logic"
	rewardschargelimittypes "fgame/fgame/game/welfare/rewards/charge_limit/types"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeRewards, welfaretypes.OpenActivityRewardsSubTypeChargeLimit, playerwelfare.ActivityObjInfoRefreshHandlerFunc(rewardsChargeRefreshInfo))
}

// 充值返利(全服次数)-刷新
func rewardsChargeRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	info := obj.GetActivityData().(*rewardschargelimittypes.ChargeRewLimitInfo)
	isFix := info.Fix()

	if isFix {
		chargelimitlogic.CheckRewardsMail(obj)
	}

	return
}
