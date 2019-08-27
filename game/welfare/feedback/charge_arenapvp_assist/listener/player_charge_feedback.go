package listener

// import (
// 	"fgame/fgame/core/event"
// 	chargeeventtypes "fgame/fgame/game/charge/event/types"
// 	gameevent "fgame/fgame/game/event"
// 	"fgame/fgame/game/player"
// 	playertypes "fgame/fgame/game/player/types"
// 	feedbackchargetypes "fgame/fgame/game/welfare/feedback/charge/types"
// 	welfarelogic "fgame/fgame/game/welfare/logic"
// 	"fgame/fgame/game/welfare/pbutil"
// 	playerwelfare "fgame/fgame/game/welfare/player"
// 	welfaretemplate "fgame/fgame/game/welfare/template"
// 	welfaretypes "fgame/fgame/game/welfare/types"
// )

// func playerChargeArenapvpAssistFeedback(target event.EventTarget, data event.EventData) (err error) {
// 	pl, ok := target.(player.Player)
// 	if !ok {
// 		return
// 	}
// 	goldNum, ok := data.(int32)
// 	if !ok {
// 		return
// 	}
// 	typ := welfaretypes.OpenActivityTypeFeedback
// 	subType := welfaretypes.OpenActivityFeedbackSubTypeCharge

// 	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
// 	feedbackTimeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
// 	for _, timeTemp := range feedbackTimeTempList {
// 		groupId := timeTemp.Group
// 		if !welfarelogic.IsOnActivityTime(groupId) {
// 			continue
// 		}

// 		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
// 		info, ok := obj.GetActivityData().(*feedbackchargetypes.FeedbackChargeInfo)
// 		if !ok {
// 			continue
// 		}

// 		info.GoldNum += goldNum
// 		welfareManager.UpdateObj(obj)

// 		scMsg := pbutil.BuildSCOpenActivityFeedbackChargeNotice(groupId, int64(info.GoldNum))
// 		pl.SendMsg(scMsg)
// 	}

// 	return
// }

// func init() {
// 	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeChargeGold, event.EventListenerFunc(playerChargeArenapvpAssistFeedback))
// }
