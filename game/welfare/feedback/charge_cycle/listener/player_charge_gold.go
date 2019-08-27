package listener

import (
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	feedbackchargecycletypes "fgame/fgame/game/welfare/feedback/charge_cycle/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//返利-连续充值
func playerChargeFeedbackCycle(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	// goldNum, ok := data.(int32)
	// if !ok {
	// 	return
	// }
	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeCycleCharge

	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		//连续充值活动
		welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
		err = welfareManager.RefreshActivityDataByGroupId(groupId)
		if err != nil {
			return
		}

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*feedbackchargecycletypes.FeedbackCycleChargeInfo)
		// groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		// if groupInterface == nil {
		// 	continue
		// }
		// groupTemp := groupInterface.(*feedbackchargecycletemplate.GroupTemplateCycleCharge)
		// needGold := groupTemp.GetDayRewCondition(info.CycleDay)

		// info.AddDayCharge(goldNum, needGold)
		// welfareManager.UpdateObj(obj)

		scMsg := pbutil.BuildSCOpenActivityFeedbackCycleChargeNotice(info.CurDayChargeNum)
		pl.SendMsg(scMsg)

	}
	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeFeedbackCycle))
}
