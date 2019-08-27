package types

import (
	arenapvptypes "fgame/fgame/game/arenapvp/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

type FeedbackChargeArenapvpAssistReturnInfo struct {
	CostNum  int64                      `json:"costNum"`  //消费数目
	IsEmail  bool                       `json:"isEmail"`  //是否奖励发放
	RankType arenapvptypes.ArenapvpType `json:"rankType"` //比武大会排名类型
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeAlliance, welfaretypes.OpenActivityAllianceSubTypeWuLian, (*FeedbackChargeArenapvpAssistReturnInfo)(nil))
}
