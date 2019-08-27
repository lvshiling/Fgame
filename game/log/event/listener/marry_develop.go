package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	"fgame/fgame/game/player"
	logmodel "fgame/fgame/logserver/model"
)

//玩家表白等级日志
func playerMarryDevelopLog(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*marryeventtypes.PlayerDevelopLevelLogEventData)
	if !ok {
		return
	}

	logMarryDevelop := &logmodel.PlayerMarryDevelop{}
	logMarryDevelop.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logMarryDevelop.BeforeLevel = eventData.GetBeforeDevelopLevel()
	logMarryDevelop.CurLevel = eventData.GetCurDevelopLevel()
	logMarryDevelop.ChangedLevel = eventData.GetCurDevelopLevel() - eventData.GetBeforeDevelopLevel()
	logMarryDevelop.Reason = int32(eventData.GetReason())
	logMarryDevelop.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logMarryDevelop)
	return
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypePlayerMarryDevelopLog, event.EventListenerFunc(playerMarryDevelopLog))
}
