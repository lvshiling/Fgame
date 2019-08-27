package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	wardrobeeventtypes "fgame/fgame/game/wardrobe/event/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家衣橱套装等级日志
func playerWardrobePeiYangLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*wardrobeeventtypes.WardrobePeiYangLogEventData)
	if !ok {
		return
	}

	logWardrobe := &logmodel.PlayerWardrobe{}
	logWardrobe.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logWardrobe.BeforeLevel = eventData.GetBeforeLev()
	logWardrobe.CurLevel = eventData.GetCurLevel()
	logWardrobe.Type = int32(eventData.GetType())
	logWardrobe.Reason = int32(eventData.GetReason())
	logWardrobe.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logWardrobe)
	return
}

func init() {
	gameevent.AddEventListener(wardrobeeventtypes.EventTypeWardrobePeiYangLog, event.EventListenerFunc(playerWardrobePeiYangLog))
}
