package player

import (
	chargesingleallmultipletypes "fgame/fgame/game/welfare/cycle/charge_single_all_multiple/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeCycleCharge, welfaretypes.OpenActivityCycleChargeSubTypeSingleChargeAllRew, playerwelfare.ActivityObjInfoInitFunc(initInfo))
}

func initInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*chargesingleallmultipletypes.CycleChargeSingleAllMultipleInfo)
	info.SingleChargeRecord = make([]int32, 1)
	info.CycleDay = welfarelogic.CountCycleDay(groupId)
	info.CanRewRecord = make(map[int32]int32)
	info.NewSingleChargeRecord = make(map[int32]int32)
	info.RewRecord = make(map[int32]int32)
}
