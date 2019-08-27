package listener

import (
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	feedbackchargereturnleveltemplate "fgame/fgame/game/welfare/feedback/charge_return_level/template"
	feedbackchargereturnleveltypes "fgame/fgame/game/welfare/feedback/charge_return_level/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家充值元宝
func playerChargeReturn(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	chargeGold, ok := data.(int32)
	if !ok {
		return
	}

	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeChargeReturnLevel
	//充值返还
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*feedbackchargereturnleveltypes.FeedbackChargeReturnLevelInfo)
		if info.IsReturn {
			continue
		}
		info.IsReturn = true
		welfareManager.UpdateObj(obj)

		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			continue
		}
		groupTemp, ok := groupInterface.(*feedbackchargereturnleveltemplate.GroupTemplateChargeReturnLevel)
		if !ok {
			continue
		}
		welfarelogic.AddExtralRewGold(pl, chargeGold, groupTemp.GetReturnType(), groupTemp.GetReturnRatio(chargeGold))
	}

	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeReturn))
}
