package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	feedbackchargedeveloptemplate "fgame/fgame/game/welfare/feedback/charge_develop/template"
	feedbackchargedeveloptypes "fgame/fgame/game/welfare/feedback/charge_develop/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeChargeDevelop, reddot.HandlerFunc(handleRedDotOpenFeedbackDevelop))
}

//养鸡生金蛋 红点
func handleRedDotOpenFeedbackDevelop(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {
	groupId := obj.GetGroupId()

	info := obj.GetActivityData().(*feedbackchargedeveloptypes.FeedbackDevelopInfo)
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*feedbackchargedeveloptemplate.GroupTemplateDevelop)
	// // 激活条件
	// if !info.IsActivate {
	// 	condition := groupTemp.GetReviveNeedGold()
	// 	if info.ActivateChargeNum >= condition {
	// 		isNotice = true
	// 		return
	// 	}
	// }

	// 喂养条件
	if !info.IsDead && !info.IsFeed {
		needCost := groupTemp.GetDevelopFeedCondition(info.FeedTimes)
		if info.TodayCostNum >= needCost {
			isNotice = true
			return
		}
	}

	// 最终奖励条件
	if !info.IsDead && !info.IsFeed {
		needFeedTimes := groupTemp.GetDevelopNeedTotalTimes()
		if info.FeedTimes >= needFeedTimes {
			isNotice = true
			return
		}
	}

	// //复活条件
	// if info.IsDead {
	// 	condition := groupTemp.GetReviveNeedGold()
	// 	if info.ActivateChargeNum >= condition {
	// 		isNotice = true
	// 		return
	// 	}
	// }

	return
}
