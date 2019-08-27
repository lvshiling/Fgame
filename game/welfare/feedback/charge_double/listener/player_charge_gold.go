package listener

import (
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	feedbackchargedoubletemplate "fgame/fgame/game/welfare/feedback/charge_double/template"
	feedbackchargedoubletypes "fgame/fgame/game/welfare/feedback/charge_double/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家充值元宝
func playerChargeDouble(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*chargeeventtypes.PlayerChargeSuccessEventData)
	if !ok {
		return
	}
	chargeGold := eventData.GetChargeGold()
	chargeId := eventData.GetChargeId()

	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeChargeDouble
	//充值翻倍
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*feedbackchargedoubletypes.FeedbackChargeDoubleInfo)
		if info.IsDouble(chargeId) {
			continue
		}

		info.AddRecord(chargeId)
		welfareManager.UpdateObj(obj)

		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			continue
		}
		groupTemp, ok := groupInterface.(*feedbackchargedoubletemplate.GroupTemplateChargeDouble)
		if !ok {
			continue
		}

		welfarelogic.AddExtralRewGold(pl, chargeGold, groupTemp.GetReturnType(), groupTemp.GetReturnRatio())
	}

	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeSuccess, event.EventListenerFunc(playerChargeDouble))
}
