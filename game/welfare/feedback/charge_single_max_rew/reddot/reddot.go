package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	feedbackchargesinglemaxrewtypes "fgame/fgame/game/welfare/feedback/charge_single_max_rew/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeSingleChagreMaxRew, reddot.HandlerFunc(handleRedDotSingleChargeMaxRew))
}

//单笔充值红点
func handleRedDotSingleChargeMaxRew(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	info := obj.GetActivityData().(*feedbackchargesinglemaxrewtypes.FeedbackSingleChargeMaxRewInfo)
	if len(info.CanRewRecord) > 0 {
		isNotice = true
		return
	}

	return
}
