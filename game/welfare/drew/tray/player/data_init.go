package player

import (
	drewtraytypes "fgame/fgame/game/welfare/drew/tray/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 幸运大转盘
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeTray, playerwelfare.ActivityObjInfoInitFunc(luckyDrewInitInfo))
}

func luckyDrewInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*drewtraytypes.LuckyDrewInfo)
	info.AttendTimes = 0
}
