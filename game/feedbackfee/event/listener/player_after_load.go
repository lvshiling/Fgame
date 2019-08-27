package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/center/center"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/feedbackfee/feedbackfee"
	feedbackfeelogic "fgame/fgame/game/feedbackfee/logic"
	"fgame/fgame/game/feedbackfee/pbutil"
	playerfeedbackfee "fgame/fgame/game/feedbackfee/player"
	feedbackfeetypes "fgame/fgame/game/feedbackfee/types"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	playertypes "fgame/fgame/game/player/types"
)

// 玩家加载后
func playerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	//下发逆付费信息
	feedbackfeeManager := pl.GetPlayerDataManager(playertypes.PlayerFeedbackFeeDataManagerType).(*playerfeedbackfee.PlayerFeedbackFeeManager)
	feeInfo := feedbackfeeManager.GetFeedbackFeeInfo()
	recordObj := feedbackfeeManager.GetCurrentRecord()
	xianJinflag := center.GetCenterService().IsXianJinOpen()
	//发送通知
	scFeedbackFeeInfo := pbutil.BuildSCFeedbackFeeInfo(xianJinflag, feeInfo, recordObj)
	pl.SendMsg(scFeedbackFeeInfo)

	//重新上传没生成code
	initRecordList := feedbackfeeManager.GetInitRecordList()
	for _, record := range initRecordList {
		feedbackfee.GetFeedbackFeeService().Exchange(pl.GetId(), record.GetId(), record.GetMoney(), record.GetExpiredTime())
	}

	generateList := feedbackfee.GetFeedbackFeeService().GetCodeGenerateList(pl.GetId())
	for _, generateCodeObj := range generateList {
		feedbackfeelogic.PlayerCodeGenerate(pl, generateCodeObj)
	}
	endList := feedbackfee.GetFeedbackFeeService().GetEndList(pl.GetId())
	for _, endObj := range endList {
		if endObj.GetStatus() == feedbackfeetypes.FeedbackExchangeStatusFailed {
			feedbackfeelogic.PlayerCodeExpire(pl, endObj)
			continue
		}
		if endObj.GetStatus() == feedbackfeetypes.FeedbackExchangeStatusFinish {
			feedbackfeelogic.PlayerCodeFinish(pl, endObj)
			continue
		}

	}

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoad))
}
