package logic

import (
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	feedbackhouseextendedtypes "fgame/fgame/game/welfare/feedback/house_extended/types"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

func SendHouseExtendedInfo(pl player.Player, groupId int32) {
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	welfareManager.RefreshActivityDataByGroupId(groupId)
	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeHouseExtended

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	obj := welfareManager.GetOpenActivityIfNotCreate(typ, subType, groupId)
	var record []int32
	activateChargeNum := int32(0)
	uplevelChargeNum := int32(0)
	uplevelGiftLevel := int32(0)
	isActivateGift := false
	isUplevelGift := false
	if obj != nil {
		info := obj.GetActivityData().(*feedbackhouseextendedtypes.FeedbackHouseExtendedInfo)
		activateChargeNum = info.ActivateChargeNum
		uplevelChargeNum = info.UplevelChargeNum
		uplevelGiftLevel = info.CurUplevelGiftLevel
		isActivateGift = info.IsActivateGift
		isUplevelGift = info.IsUplevelGift
	}

	scMsg := pbutil.BuildSCOpenActivityGetInfoHouseExtended(groupId, startTime, endTime, record, activateChargeNum, uplevelChargeNum, uplevelGiftLevel, isActivateGift, isUplevelGift)
	pl.SendMsg(scMsg)
}
