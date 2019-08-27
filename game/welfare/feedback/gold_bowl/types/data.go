package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//返利-聚宝盆
type FeedbackGoldBowlInfo struct {
	GoldNum int64 `json:"goldNum"` //花费元宝数量
	IsEmail bool  `json:"isEmail"` //是否奖励发放
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeGoldBowl, (*FeedbackGoldBowlInfo)(nil))
}
