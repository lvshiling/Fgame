package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家拉霸日志
func playerAttendLaBaLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*welfareeventtypes.PlayerLaBaLogEventData)
	if !ok {
		return
	}

	logLaba := &logmodel.PlayerLaBa{}
	logLaba.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logLaba.CurTimes = eventData.GetAttendTimes()
	logLaba.CostGold = eventData.GetCostGold()
	logLaba.RewGold = eventData.GetRewGold()
	logLaba.Reason = int32(eventData.GetReason())
	logLaba.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logLaba)
	return
}

func init() {
	gameevent.AddEventListener(welfareeventtypes.EventTypeLaBaAttendLog, event.EventListenerFunc(playerAttendLaBaLog))
}
