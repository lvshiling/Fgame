package listener

import (
	"fgame/fgame/core/event"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	newchargearenapvpvassitreturntypes "fgame/fgame/game/welfare/alliance/new_charge_arenapvp_assist_return/types"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	"fgame/fgame/game/welfare/welfare"
)

//购买转生大礼包
func discountBuyZhuanShengGift(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*welfareeventtypes.PlayerAllianceCheerEventData)
	if !ok {
		return
	}
	groupId := eventData.GetGroupId()
	costGold := eventData.GetCostGold()
	giftType := eventData.GetGiftType()

	//城战助威礼包
	buyNewChargeArenapvpAssistGift(pl, groupId, costGold, giftType)
	return
}

func buyNewChargeArenapvpAssistGift(pl player.Player, groupId, costGold, giftType int32) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	for _, relateGroupId := range timeTemp.GetRelationToGroupList() {
		if !welfarelogic.IsOnActivityTime(relateGroupId) {
			continue
		}

		//城战助威
		relateTimeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(relateGroupId)
		if relateTimeTemp.IsArenaPvpCheer() {
			activityTemp := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeArenapvp)
			if activityTemp == nil {
				continue
			}
			// 是否活动日
			now := global.GetGame().GetTimeService().Now()
			openTime := global.GetGame().GetServerTime()
			mergeTime := merge.GetMergeService().GetMergeTime()
			activityTimeTemp := activityTemp.GetOnDateTimeTemplate(now, openTime, mergeTime)
			if activityTimeTemp == nil {
				continue
			}

			// 是否活动开始
			beginTime, _ := activityTimeTemp.GetBeginTime(now)
			if now >= beginTime {
				continue
			}

			relateObj := welfareManager.GetOpenActivityIfNotCreate(relateTimeTemp.GetOpenType(), relateTimeTemp.GetOpenSubType(), relateGroupId)
			info := relateObj.GetActivityData().(*newchargearenapvpvassitreturntypes.FeedbackNewChargeArenapvpAssistReturnInfo)
			info.AddCostNum(int64(costGold))
			welfareManager.UpdateObj(relateObj)
			startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(relateGroupId)
			scMsg := pbutil.BuildSCOpenActivityFeedbackChargeNewArenapvpAssistReturnInfo(relateGroupId, startTime, endTime, info.CostNum)
			pl.SendMsg(scMsg)
		}

	}
}

func init() {
	gameevent.AddEventListener(welfareeventtypes.EventTypeDiscountBuyZhuanShengGift, event.EventListenerFunc(discountBuyZhuanShengGift))
}
