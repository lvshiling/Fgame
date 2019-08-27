package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	playerproperty "fgame/fgame/game/property/player"
	feedbackchargedeveloptypes "fgame/fgame/game/welfare/feedback/charge_develop/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家消耗元宝
func playerCostGoldFeedbackDevelop(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	// goldNum, ok := data.(int64)
	// if !ok {
	// 	return
	// }
	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeChargeDevelop

	//金鸡培养
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		// 刷新
		err = welfareManager.RefreshActivityDataByGroupId(groupId)
		if err != nil {
			return
		}

		obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
		info := obj.GetActivityData().(*feedbackchargedeveloptypes.FeedbackDevelopInfo)
		if !info.IsActivate || info.IsDead {
			continue
		}

		info.TodayCostNum = propertyManager.GetTodayCostNum()
		welfareManager.UpdateObj(obj)

	}

	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerGoldCost, event.EventListenerFunc(playerCostGoldFeedbackDevelop))
}
