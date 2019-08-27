package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//返利-养金猪
type FeedbackGoldPigInfo struct {
	CurCondition int32 `json:"curCondition"` //当前充值元宝条件
	ChargeGold   int32 `json:"chargeGold"`   //充值元宝数量
	CostGold     int32 `json:"costGold"`     //花费元宝数量
	IsEmail      bool  `json:"isEmail"`      //是否奖励发放
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeGoldPig, (*FeedbackGoldPigInfo)(nil))
}
