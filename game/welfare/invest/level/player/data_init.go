package player

import (
	investleveltypes "fgame/fgame/game/welfare/invest/level/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 投资计划
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeInvest, welfaretypes.OpenActivityInvestSubTypeLevel, playerwelfare.ActivityObjInfoInitFunc(investLevelInitInfo))
}

func investLevelInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*investleveltypes.InvestLevelInfo)
	info.InvestBuyInfoMap = make(map[investleveltypes.InvestLevelType]int32)
	info.IsBack = true
}
