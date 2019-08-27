package listener

import (
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	feedbackpigtemplate "fgame/fgame/game/welfare/feedback/pig/template"
	feedbackpigtypes "fgame/fgame/game/welfare/feedback/pig/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家充值元宝
func playerChargeGoldPig(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	goldNum, ok := data.(int32)
	if !ok {
		return
	}

	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeGoldPig
	//返利-元宝猪
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}
		groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
		if groupInterface == nil {
			continue
		}
		groupTemp := groupInterface.(*feedbackpigtemplate.GroupTemplateFeedbackGoldPig)

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*feedbackpigtypes.FeedbackGoldPigInfo)
		info.ChargeGold += goldNum
		info.CurCondition = groupTemp.CountGoldPigNeedCharge(info.CurCondition, info.ChargeGold)
		welfareManager.UpdateObj(obj)

		scMsg := pbutil.BuildSCOpenActivityFeedbackChargeNotice(groupId, int64(info.ChargeGold))
		pl.SendMsg(scMsg)
	}

	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeGoldPig))
}
