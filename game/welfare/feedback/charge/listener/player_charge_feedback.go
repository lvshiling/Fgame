package listener

import (
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	feedbackchargetypes "fgame/fgame/game/welfare/feedback/charge/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
)

//玩家充值元宝
func playerChargeFeedback(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	goldNum, ok := data.(int32)
	if !ok {
		return
	}
	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeCharge

	//充值返利活动
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	feedbackTimeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range feedbackTimeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		welfareManager.RefreshActivityDataByGroupId(groupId)
		feedbackInfo := obj.GetActivityData().(*feedbackchargetypes.FeedbackChargeInfo)

		// 第一天走refresh，同步今日累计充值
		now := global.GetGame().GetTimeService().Now()
		diff, _ := timeutils.DiffDay(now, obj.GetStartTime())
		if diff > 0 {
			feedbackInfo.GoldNum += goldNum
			welfareManager.UpdateObj(obj)
		}

		scMsg := pbutil.BuildSCOpenActivityFeedbackChargeNotice(groupId, int64(feedbackInfo.GoldNum))
		pl.SendMsg(scMsg)
	}

	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeFeedback))
}
