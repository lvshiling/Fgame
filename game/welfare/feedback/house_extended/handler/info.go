package handler

import (
	"fgame/fgame/game/player"
	houseextendedlogic "fgame/fgame/game/welfare/feedback/house_extended/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func init() {
	welfare.RegisterInfoGetHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeHouseExtended, welfare.InfoGetHandlerFunc(getAdvancedRewMaxInfo))
}

//获取升阶奖励请求逻辑
func getAdvancedRewMaxInfo(pl player.Player, groupId int32) (err error) {
	houseextendedlogic.SendHouseExtendedInfo(pl, groupId)
	return
}
