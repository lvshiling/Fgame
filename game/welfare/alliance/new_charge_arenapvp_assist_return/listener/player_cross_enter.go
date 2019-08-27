package listener

import (
	"fgame/fgame/core/event"
	arenapvptypes "fgame/fgame/game/arenapvp/types"
	crosseventtypes "fgame/fgame/game/cross/event/types"
	crosstypes "fgame/fgame/game/cross/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	alliancenewchargearenapvpassistreturntypes "fgame/fgame/game/welfare/alliance/new_charge_arenapvp_assist_return/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func playerCrossEnter(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	typ := welfaretypes.OpenActivityTypeAlliance
	subType := welfaretypes.OpenActivityAllianceSubTypeNewWuLian
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)

	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info, ok := obj.GetActivityData().(*alliancenewchargearenapvpassistreturntypes.FeedbackNewChargeArenapvpAssistReturnInfo)
		if !ok {
			continue
		}

		if pl.GetCrossType() == crosstypes.CrossTypeArenapvp {
			info.RankType = arenapvptypes.ArenapvpTypeElection
			// info.Rank = arenapvptypes.ArenapvpTypeElection.GetNumber()
			welfareManager.UpdateObj(obj)
		}
	}
	return
}

func init() {
	gameevent.AddEventListener(crosseventtypes.EventTypePlayerCrossEnter, event.EventListenerFunc(playerCrossEnter))
}
