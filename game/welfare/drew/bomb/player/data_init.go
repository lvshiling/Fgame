package player

import (
	drewbombtypes "fgame/fgame/game/welfare/drew/bomb/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 炸矿
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeMergeDrew, welfaretypes.OpenActivityDrewSubTypeBombOre, playerwelfare.ActivityObjInfoInitFunc(bombOreInitInfo))
}

func bombOreInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*drewbombtypes.BombOreInfo)
	info.AttendTimes = 0
}
