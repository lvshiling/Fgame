package reddot

import (
	"fgame/fgame/game/player"
	"fgame/fgame/game/reddot/reddot"
	feedbackhouseextendedtemplate "fgame/fgame/game/welfare/feedback/house_extended/template"
	feedbackhouseextendedtypes "fgame/fgame/game/welfare/feedback/house_extended/types"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
)

func init() {
	reddot.Register(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeHouseExtended, reddot.HandlerFunc(handleRedDotOpenHouseExtended))
}

//开服-房产活动红点
func handleRedDotOpenHouseExtended(pl player.Player, obj *playerwelfare.PlayerOpenActivityObject) (isNotice bool) {

	groupId := obj.GetGroupId()
	info := obj.GetActivityData().(*feedbackhouseextendedtypes.FeedbackHouseExtendedInfo)
	groupInterface := welfaretemplate.GetWelfareTemplateService().GetOpenActivityGroupTemplateInterface(groupId)
	if groupInterface == nil {
		return
	}
	groupTemp := groupInterface.(*feedbackhouseextendedtemplate.GroupTemplateHouseExtended)

	// 激活礼包
	if !info.IsActivateGift {
		if groupTemp.GetActivateCanRewTemp(info.ActivateChargeNum) == nil {
			return
		}

		isNotice = true
		return
	}

	//升级礼包
	if !info.IsUplevelGift {
		if groupTemp.GetUplevelCanRewTemp(info.UplevelChargeNum, info.CurUplevelGiftLevel) == nil {
			return
		}

		isNotice = true
		return
	}

	return
}
