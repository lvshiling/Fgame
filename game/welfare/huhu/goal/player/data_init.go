package player

import (
	goaltypes "fgame/fgame/game/welfare/huhu/goal/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 活动目标
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeHuHu, welfaretypes.OpenActivitySpecialSubTypeGoal, playerwelfare.ActivityObjInfoInitFunc(goalInitInfo))
}

func goalInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*goaltypes.GoalInfo)
	info.RewRecordMap = make(map[int32]struct{})
	info.GoalCount = 0
	info.IsEmail = false
}
