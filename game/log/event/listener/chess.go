package listener

import (
	"fgame/fgame/core/event"
	chesseventtypes "fgame/fgame/game/chess/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	logmodel "fgame/fgame/logserver/model"
)

//玩家苍龙棋局日志
func playerChessLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*chesseventtypes.PlayerAttendChessLogEventData)
	if !ok {
		return
	}

	logChess := &logmodel.PlayerChess{}
	logChess.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logChess.AttendTimes = eventData.GetAttendTimes()
	logChess.Reason = int32(eventData.GetReason())
	logChess.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logChess)
	return
}

func init() {
	gameevent.AddEventListener(chesseventtypes.EventTypeAttendChessLog, event.EventListenerFunc(playerChessLog))
}
