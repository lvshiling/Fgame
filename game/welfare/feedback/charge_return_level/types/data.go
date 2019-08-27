package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//首充返还(活动时间内享受一次)
type FeedbackChargeReturnLevelInfo struct {
	IsReturn bool `json:"isReturn"` //是否享受返还
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeReturnLevel, (*FeedbackChargeReturnLevelInfo)(nil))
}
