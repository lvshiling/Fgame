package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
	feedbackchargereturnmultipletypes "fgame/fgame/game/welfare/feedback/charge_return_multiple/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretemplate "fgame/fgame/game/welfare/template"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"
)

//加载完成后
func playerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	typ := welfaretypes.OpenActivityTypeFeedback
	subType := welfaretypes.OpenActivityFeedbackSubTypeChargeReturnMultiple
	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)

	timeTempList := welfaretemplate.GetWelfareTemplateService().GetOpenActivityTimeTemplateByType(typ, subType)
	for _, timeTemp := range timeTempList {
		groupId := timeTemp.Group
		if !welfarelogic.IsOnActivityTime(groupId) {
			continue
		}

		startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
		obj := welfareManager.GetOpenActivity(groupId)
		periodChargeNum := int32(0)
		rewardCnt := int32(0)
		var record []int32
		if obj != nil {
			err = welfareManager.RefreshActivityDataByGroupId(groupId)
			if err != nil {
				return
			}
			info := obj.GetActivityData().(*feedbackchargereturnmultipletypes.FeedbackChargeReturnMultipleInfo)
			periodChargeNum = info.PeriodChargeNum
			rewardCnt = info.RewardCnt
		}

		scMsg := pbutil.BuildSCOpenActivityGetInfoChargeReturnMultiple(groupId, startTime, endTime, record, periodChargeNum, rewardCnt)
		pl.SendMsg(scMsg)
	}

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoad))
}
