package player

import (
	playercharge "fgame/fgame/game/charge/player"
	"fgame/fgame/game/global"
	playertypes "fgame/fgame/game/player/types"
	feedbacklabatypes "fgame/fgame/game/welfare/feedback/laba/types"
	welfarelogic "fgame/fgame/game/welfare/logic"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/pkg/timeutils"
)

func init() {
	playerwelfare.RegisterInfoRefreshHandler(welfaretypes.OpenActivityTypeFeedback, welfaretypes.OpenActivityFeedbackSubTypeGoldLaBa, playerwelfare.ActivityObjInfoRefreshHandlerFunc(feedbackLaBaRefreshInfo))
}

//元宝拉霸-刷新
func feedbackLaBaRefreshInfo(obj *playerwelfare.PlayerOpenActivityObject) (err error) {
	// 同步今日充值
	if welfarelogic.IsOnActivityTime(obj.GetGroupId()) {
		now := global.GetGame().GetTimeService().Now()
		diff, _ := timeutils.DiffDay(now, obj.GetStartTime())
		if diff == 0 {
			pl := obj.GetPlayer()
			welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
			chargeManager := pl.GetPlayerDataManager(playertypes.PlayerChargeDataManagerType).(*playercharge.PlayerChargeDataManager)
			info := obj.GetActivityData().(*feedbacklabatypes.FeedbackGoldLaBaInfo)
			if info.ChargeNum != int32(chargeManager.GetTodayChargeNum()) {
				info.ChargeNum = int32(chargeManager.GetTodayChargeNum())
				welfareManager.UpdateObj(obj)
			}
		}
	}

	return
}
