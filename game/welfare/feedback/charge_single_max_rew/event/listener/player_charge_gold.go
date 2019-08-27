package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/core/utils"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	feedbackchargesinglemaxrewtemplate "fgame/fgame/game/welfare/feedback/charge_single_max_rew/template"
	feedbackchargesinglemaxrewtypes "fgame/fgame/game/welfare/feedback/charge_single_max_rew/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//单笔充值活动(最近档次)
func playerChargeMergeSingle(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	goldNum, ok := data.(int32)
	if !ok {
		return
	}
	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeSingleChagreMaxRew

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*feedbackchargesinglemaxrewtypes.FeedbackSingleChargeMaxRewInfo)
		if goldNum > info.MaxSingleChargeNum {
			info.MaxSingleChargeNum = goldNum
		}
		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			continue
		}
		groupTemp := groupInterface.(*feedbackchargesinglemaxrewtemplate.GroupTemplateSingleChargeMaxRew)
		descTempList := groupTemp.GetTempDescList()
		for _, temp := range descTempList {
			needGold := temp.Value1
			if goldNum < needGold {
				continue
			}

			if utils.ContainInt32(info.CanRewRecord, needGold) {
				continue
			}

			if utils.ContainInt32(info.ReceiveRewRecord, needGold) {
				continue
			}

			info.AddCanRewRecord(needGold)
			break
		}

		welfareManager.UpdateObj(obj)

		scMsg := pbutil.BuildSCOpenActivityCycleSingleChargeMaxRewInfoNotice(groupId, info.MaxSingleChargeNum, info.CanRewRecord)
		pl.SendMsg(scMsg)

	}
	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeMergeSingle))
}
