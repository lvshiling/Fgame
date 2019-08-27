package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 返利-元宝拉霸
type FeedbackGoldLaBaInfo struct {
	Times     int32 `json:"times"`     //拉霸次数
	ChargeNum int32 `json:"chargeNum"` //当前充值金额
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeGoldLaBa, (*FeedbackGoldLaBaInfo)(nil))
}
