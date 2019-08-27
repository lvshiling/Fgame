package listener

import (
	"fgame/fgame/core/event"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	alliancecheertypes "fgame/fgame/game/welfare/alliance/cheer/types"
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
	buyAllianceCheerGift(pl, groupId, costGold, giftType)
	return
}

func buyAllianceCheerGift(pl player.Player, groupId, costGold, giftType int32) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(groupId)
	for _, relateGroupId := range timeTemp.GetRelationToGroupList() {
		if !welfarelogic.IsOnActivityTime(relateGroupId) {
			continue
		}

		//城战助威
		relateTimeTemp := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplate(relateGroupId)
		if relateTimeTemp.IsAllianceCheer() {
			activityTemp := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeAlliance)
			if activityTemp == nil {
				continue
			}
			// 是否活动日
			now := global.GetGame().GetTimeService().Now()
			openTime := welfare.GetWelfareService().GetServerStartTime()  //global.GetGame().GetServerTime()
			mergeTime := welfare.GetWelfareService().GetServerMergeTime() //merge.GetMergeService().GetMergeTime()
			activityTimeTemp := activityTemp.GetOnDateTimeTemplate(now, openTime, mergeTime)
			if activityTimeTemp == nil {
				continue
			}

			// 是否活动结束
			endTime, _ := activityTimeTemp.GetEndTime(now)
			if now >= endTime {
				continue
			}

			relateObj := welfareManager.GetOpenActivityIfNotCreate(relateTimeTemp.GetOpenType(), relateTimeTemp.GetOpenSubType(), relateGroupId)
			info := relateObj.GetActivityData().(*alliancecheertypes.AllianceCheerInfo)
			info.CheerGoldNum += costGold
			welfareManager.UpdateObj(relateObj)

			scMsg := pbutil.BuildSCOpenActivityAllianceCheerPoolChanged(groupId, int64(info.CheerGoldNum))
			pl.SendMsg(scMsg)
		}

	}
}

func init() {
	gameevent.AddEventListener(welfareeventtypes.EventTypeDiscountBuyZhuanShengGift, event.EventListenerFunc(discountBuyZhuanShengGift))
}
