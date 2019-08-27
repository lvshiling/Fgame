package player

import (
	feedbacknewchargearenapvpassistreturntypes "fgame/fgame/game/welfare/alliance/new_charge_arenapvp_assist_return/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	playerwelfare.RegisterInfoInitHandler(welfaretypes.OpenActivityTypeAlliance, welfaretypes.OpenActivityAllianceSubTypeNewWuLian, playerwelfare.ActivityObjInfoInitFunc(feedbackChargeNewArenapvpAssistReturnInitInfo))
}

func feedbackChargeNewArenapvpAssistReturnInitInfo(obj *playerwelfare.PlayerOpenActivityObject) {
	info := obj.GetActivityData().(*feedbacknewchargearenapvpassistreturntypes.FeedbackNewChargeArenapvpAssistReturnInfo)

	info.Reset()
}
