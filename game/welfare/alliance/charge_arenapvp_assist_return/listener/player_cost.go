package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	alliancechargearenapvpassistreturntypes "fgame/fgame/game/welfare/alliance/charge_arenapvp_assist_return/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
)

func playerCostGoldFeedbackDevelop(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	goldNum, ok := data.(int64)
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
		welfareManager.RefreshActivityDataByGroupId(groupId)
		info, ok := obj.GetActivityData().(*alliancechargearenapvpassistreturntypes.FeedbackChargeArenapvpAssistReturnInfo)
		if !ok {
			continue
		}

		// 第一天走refresh，同步今日消费
		now := global.GetGame().GetTimeService().Now()
		diff, _ := timeutils.DiffDay(now, obj.GetStartTime())
		if diff > 0 {
			info.CostNum += goldNum
			welfareManager.UpdateObj(obj)
		}

		scMsg := pbutil.BuildSCOpenActivityFeedbackCostNotice(groupId, int32(info.CostNum))
		pl.SendMsg(scMsg)
	}
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerGoldCost, event.EventListenerFunc(playerCostGoldFeedbackDevelop))
}
