package player

import (
	investsevendaytypes "fgame/fgame/game/welfare/invest/sevenday/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 七日投资
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeInvest, welfaretypes.OpenActivityInvestSubTypeServenDay, playerwelfare.ActivityObjInfoInitFunc(investDayInitInfo))
}

func investDayInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*investsevendaytypes.InvestDayInfo)
	info.ReceiveDay = 0
	info.BuyTime = 0
}
