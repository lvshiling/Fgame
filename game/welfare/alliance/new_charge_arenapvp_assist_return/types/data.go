package types

import (
	arenapvptypes "fgame/fgame/game/arenapvp/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

type FeedbackNewChargeArenapvpAssistReturnInfo struct {
	CostNum  int64                      `json:"costNum"`  //消费数目
	IsEmail  bool                       `json:"isEmail"`  //是否奖励发放
	RankType arenapvptypes.ArenapvpType `json:"rankType"` //比武大会排名类型
}

func (info *FeedbackNewChargeArenapvpAssistReturnInfo) AddCostNum(costNum int64) {
	info.CostNum += costNum
}

func (info *FeedbackNewChargeArenapvpAssistReturnInfo) UpdateRankType(rankType arenapvptypes.ArenapvpType) {
	info.RankType = rankType
}

func (info *FeedbackNewChargeArenapvpAssistReturnInfo) Email() {
	info.IsEmail = true
}

func (info *FeedbackNewChargeArenapvpAssistReturnInfo) Reset() {
	info.IsEmail = false
	info.CostNum = 0
	info.RankType = arenapvptypes.ArenapvpTypeInit
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeAlliance, welfaretypes.OpenActivityAllianceSubTypeNewWuLian, (*FeedbackNewChargeArenapvpAssistReturnInfo)(nil))
}
