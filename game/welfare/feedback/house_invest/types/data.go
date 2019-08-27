package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//返利-房产投资
type FeedbackHouseInvestInfo struct {
	ChargeNum       int32 `json:"chargeNum"`       //活动区间充值数量
	IsActivity      bool  `json:"isActivity"`      //是否激活
	CurDayChargeNum int32 `json:"curDayChargeNum"` //今日充值数量
	IsCurDayDecor   bool  `json:"isCurDayDecor"`   //今天是否装修
	DecorDays       int32 `json:"decorDays"`       //装修几天
	IsSell          bool  `json:"isSell"`          //是否出售
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeHouseInvest, (*FeedbackHouseInvestInfo)(nil))
}
