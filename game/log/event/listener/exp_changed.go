package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	logmodel "fgame/fgame/logserver/model"
)

//经验变化
func playerExpChanged(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*propertyeventtypes.PlayerExpChangedLogEventData)
	if !ok {
		return
	}

	logExpChanged := &logmodel.PlayerExpChanged{}
	logExpChanged.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logExpChanged.BeforeExp = eventData.GetBeforeExp()
	logExpChanged.CurExp = eventData.GetCurExp()
	logExpChanged.BeforeLevel = eventData.GetBeforeLevel()
	logExpChanged.CurLevel = eventData.GetCurLevel()
	logExpChanged.ChangedExp = eventData.GetChangedExp()
	logExpChanged.Reason = eventData.GetReason().Reason()
	logExpChanged.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logExpChanged)
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerExpChangedLog, event.EventListenerFunc(playerExpChanged))
}
