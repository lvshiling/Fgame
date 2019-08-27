package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	feedbackeventtypes "fgame/fgame/game/feedbackfee/event/types"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	logmodel "fgame/fgame/logserver/model"
)

//逆付费兑换日志
func feedbackfeeExchangeLog(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*feedbackeventtypes.FeedbackfeeExchangeLogEventData)
	if !ok {
		return
	}

	logFeedbackfee := &logmodel.PlayerFeedbackfee{}
	logFeedbackfee.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logFeedbackfee.CurMoney = eventData.GetCurMoney()
	logFeedbackfee.BeforeMoney = eventData.GetBeforeMoney()
	logFeedbackfee.Changed = eventData.GetChanged()
	logFeedbackfee.Reason = int32(eventData.GetReason())
	logFeedbackfee.ReasonText = eventData.GetReasonText()

	log.GetLogService().SendLog(logFeedbackfee)
	return
}

func init() {
	gameevent.AddEventListener(feedbackeventtypes.EventTypeFeedbackfeeExchangeLog, event.EventListenerFunc(feedbackfeeExchangeLog))
}
