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

//玩家战力日志
func playerForceLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*propertyeventtypes.PlayerForceChangedEventData)
	if !ok {
		return
	}

	forceLog := &logmodel.PlayerForce{}
	forceLog.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	forceLog.Force = eventData.GetForce()
	forceLog.BeforeForce = eventData.GetBeforeForce()
	forceLog.Mask = eventData.GetMask()
	forceLog.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(forceLog)
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerForceChanged, event.EventListenerFunc(playerForceLog))
}
