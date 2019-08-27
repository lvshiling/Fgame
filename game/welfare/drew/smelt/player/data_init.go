package player

import (
	smelttypes "fgame/fgame/game/welfare/drew/smelt/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeSmelt, playerwelfare.ActivityObjInfoInitFunc(smeltInitInfo))
}

func smeltInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info, _ := obj.GetActivityData().(*smelttypes.SmeltInfo)
	info.Num = 0
	info.HasReceiveNum = 0
}
