package listener

import (
	"fgame/fgame/core/event"
	arenapvpeventtypes "fgame/fgame/game/arenapvp/event/types"
	playerarenapvp "fgame/fgame/game/arenapvp/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	alliancechargearenapvpassistreturntypes "fgame/fgame/game/welfare/alliance/charge_arenapvp_assist_return/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家比武大会成绩更新
func playerArenapvpResult(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	arenapvpObj, ok := data.(*playerarenapvp.PlayerArenapvpObject)
	if !ok {
		return
	}

	typ := welfaretypes.OpenActivityTypeAlliance
	subType := welfaretypes.OpenActivityAllianceSubTypeWuLian
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*alliancechargearenapvpassistreturntypes.FeedbackChargeArenapvpAssistReturnInfo)

		// rank := arenapvpobj.GetPvpRecord().GetNumber()
		// info.Rank = rank
		info.RankType = arenapvpObj.GetPvpRecord()
		welfareManager.UpdateObj(obj)
	}
	return
}

func init() {
	gameevent.AddEventListener(arenapvpeventtypes.EventTypeArenapvpResult, event.EventListenerFunc(playerArenapvpResult))
}
