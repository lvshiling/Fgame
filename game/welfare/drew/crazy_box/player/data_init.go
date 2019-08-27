package player

import (
	drewcrazyboxtypes "fgame/fgame/game/welfare/drew/crazy_box/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 疯狂宝箱
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeCrazyBox, playerwelfare.ActivityObjInfoInitFunc(crazyBoxInitInfo))
}

func crazyBoxInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*drewcrazyboxtypes.CrazyBoxInfo)
	info.GoldNum = 0
	info.AttendTimes = 0
}
