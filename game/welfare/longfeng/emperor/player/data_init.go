package player

import (
	longfengemperortypes "fgame/fgame/game/welfare/longfeng/emperor/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 龙凤呈祥
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeLongFeng, welfaretypes.OpenActivityDefaultSubTypeDefault, playerwelfare.ActivityObjInfoInitFunc(longFengInitInfo))
}

func longFengInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*longfengemperortypes.LongFengInfo)
	info.RobTimes = 0
}
