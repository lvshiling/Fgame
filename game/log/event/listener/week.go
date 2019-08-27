package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	weekeventtypes "fgame/fgame/game/week/event/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家周卡购买日志
func playerWeekBuyLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*weekeventtypes.PlayerWeekLogBuyLogEventData)
	if !ok {
		return
	}

	weekLog := &logmodel.PlayerWeek{}
	weekLog.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	weekLog.LastExpireTime = eventData.GetLastExpireTime()
	weekLog.WeekType = eventData.GetWeekType()
	weekLog.Reason = int32(eventData.GetReason())
	weekLog.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(weekLog)
	return
}

func init() {
	gameevent.AddEventListener(weekeventtypes.EventTypeWeekLogBuy, event.EventListenerFunc(playerWeekBuyLog))
}
