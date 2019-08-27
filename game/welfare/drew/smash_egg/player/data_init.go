package player

import (
	drewsmasheggtypes "fgame/fgame/game/welfare/drew/smash_egg/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 砸金蛋
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeSmashEgg, playerwelfare.ActivityObjInfoInitFunc(smashEggInitInfo))
}

func smashEggInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*drewsmasheggtypes.SmashEggInfo)
	info.AttendTimes = 0
	info.AttendTimesList = []int32{}
}
