package player

import (
	investnewleveltypes "fgame/fgame/game/welfare/invest/new_level/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 新等级投资计划
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeInvest, welfaretypes.OpenActivityInvestSubTypeNewLevel, playerwelfare.ActivityObjInfoInitFunc(investNewLevelInitInfo))
}

func investNewLevelInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*investnewleveltypes.InvestNewLevelInfo)
	info.InvestBuyInfoMap = make(map[investnewleveltypes.InvestNewLevelType][]int32)
}
