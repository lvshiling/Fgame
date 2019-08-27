package player

import (
	advancedfeedbacktypes "fgame/fgame/game/welfare/advanced/feedback/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

// 进阶返利（废弃）
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeAdvanced, welfaretypes.OpenActivityAdvancedSubTypeFeedback, playerwelfare.ActivityObjInfoInitFunc(advancedInitInfo))
}

func advancedInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*advancedfeedbacktypes.AdvancedInfo)
	info.DanNum = 0
	info.RewRecord = []int32{}
	info.AdvancedDay = welfarelogic.CountCurActivityDay(groupId)
}
