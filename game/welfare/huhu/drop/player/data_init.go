package player

import (
	huhudroptypes "fgame/fgame/game/welfare/huhu/drop/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 虎虎生风
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeHuHu, welfaretypes.OpenActivitySpecialSubTypeDrop, playerwelfare.ActivityObjInfoInitFunc(huhuInitInfo))
}

func huhuInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*huhudroptypes.HuHuInfo)
	info.CurDayDropNum = 0
}
