package player

import (
	arenapvptypes "fgame/fgame/game/arenapvp/types"
	feedbackchargearenapvpassistreturntypes "fgame/fgame/game/welfare/alliance/charge_arenapvp_assist_return/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeAlliance, welfaretypes.OpenActivityAllianceSubTypeWuLian, playerwelfare.ActivityObjInfoInitFunc(feedbackChargeArenapvpAssistReturnInitInfo))
}

func feedbackChargeArenapvpAssistReturnInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*feedbackchargearenapvpassistreturntypes.FeedbackChargeArenapvpAssistReturnInfo)
	info.CostNum = 0
	info.IsEmail = false
	info.RankType = arenapvptypes.ArenapvpType(0)
}
