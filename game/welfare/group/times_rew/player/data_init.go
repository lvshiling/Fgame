package player

import (
	grouptimesrewtypes "fgame/fgame/game/welfare/group/times_rew/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 次数奖励
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeGroup, welfaretypes.OpenActivityGroupSubTypeTimesRew, playerwelfare.ActivityObjInfoInitFunc(timesRewInitInfo))
}

func timesRewInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*grouptimesrewtypes.TimesRewInfo)
	info.RewRecord = make(map[int32][]int32)
	info.Times = 0
	info.IsEmail = false
}
