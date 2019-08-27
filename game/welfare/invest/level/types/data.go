package types

import (
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//投资-初级/高级
type InvestLevelInfo struct {
	InvestBuyInfoMap map[InvestLevelType]int32 `json:"investBuyInfo"` //投资计划购买信息
	IsBack           bool                      `json:"isBack"`        //是否返还检测过
}

func (info InvestLevelInfo) IsBuy(typ InvestLevelType) bool {
	_, isBuy := info.InvestBuyInfoMap[typ]
	return isBuy
}

func init() {
	playerwelfare.RegisterOpenActivityData(welfaretypes.OpenActivityTypeInvest, welfaretypes.OpenActivityInvestSubTypeLevel, (*InvestLevelInfo)(nil))
}
