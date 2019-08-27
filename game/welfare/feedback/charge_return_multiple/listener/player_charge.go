package listener

import (
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	playercharge "fgame/fgame/game/charge/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	feedbackchargereturnmultipletypes "fgame/fgame/game/welfare/feedback/charge_return_multiple/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
)

//累充奖励
func playerChargeReturnMultiple(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	goldNum, ok := data.(int32)
	if !ok {
		return
	}
	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeChargeReturnMultiple

	//累充奖励活动
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	err = welfareManager.RefreshActivityData(typ, subType)
	if err != nil {
		return
	}
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		now := global.GetGame().GetTimeService().Now()
		info := obj.GetActivityData().(*feedbackchargereturnmultipletypes.FeedbackChargeReturnMultipleInfo)
		diff, _ := timeutils.DiffDay(now, obj.GetStartTime())
		if diff == 0 {
			pl := obj.GetPlayer()
			chargeManager := pl.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
			welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
			if info.PeriodChargeNum != int32(chargeManager.GetTodayChargeNum()) {
				diffNum := int32(chargeManager.GetTodayChargeNum()) - info.PeriodChargeNum
				if diffNum > 0 {
					info.AddPeriodCharge(diffNum)
					welfareManager.UpdateObj(obj)
				}
			}
		} else {
			info.AddPeriodCharge(goldNum)
			welfareManager.UpdateObj(obj)
		}
		scMsg := pbutil.BuildSCOpenActivityFeedbackChargeReturnMultipleNotice(info.PeriodChargeNum)
		pl.SendMsg(scMsg)

	}
	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeReturnMultiple))
}
