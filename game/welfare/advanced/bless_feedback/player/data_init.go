package player

import (
	advancedblessfeedbacktypes "fgame/fgame/game/welfare/advanced/bless_feedback/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

// 进阶祝福丹大放送(废弃)
func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeAdvanced, welfaretypes.OpenActivityAdvancedSubTypeBlessFeedback, playerwelfare.ActivityObjInfoInitFunc(blessAdvancedInitInfo))
}

func blessAdvancedInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*advancedblessfeedbacktypes.BlessAdvancedInfo)
	info.RewRecord = []int32{}
	info.BlessDay = welfarelogic.CountCurActivityDay(groupId)

	pl := obj.GetPlayer()
	curBlessDay := welfarelogic.CountCurActivityDay(groupId)
	advancedType := welfaretypes.AdvancedType(curBlessDay)
	info.AdvancedNum = welfare.GetSystemAdvancedNum(pl, advancedType)
}
