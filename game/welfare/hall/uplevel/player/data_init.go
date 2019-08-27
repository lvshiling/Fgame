package player

import (
	hallupleveltypes "fgame/fgame/game/welfare/hall/uplevel/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 等级冲刺
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeWelfare, welfaretypes.OpenActivityWelfareSubTypeUpLevel, playerwelfare.ActivityObjInfoInitFunc(welfareLevelInitInfo))
}

func welfareLevelInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*hallupleveltypes.WelfareUplevelInfo)
	info.RewRecord = []int32{}
	info.IsEmail = false
}
