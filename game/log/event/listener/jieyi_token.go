package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	jieyieventtypes "fgame/fgame/game/jieyi/event/types"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	logmodel "fgame/fgame/logserver/model"
)

//玩家结义信物改变日志
func playerJieYiTokenLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*jieyieventtypes.PlayerJieYiTokenTypeChangeLogEventData)
	if !ok {
		return
	}

	tokenLog := &logmodel.PlayerJieYiToken{}
	tokenLog.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	tokenLog.BeforeToken = eventData.GetBeformType().String()
	tokenLog.CurToken = eventData.GetCurType().String()
	tokenLog.Reason = int32(eventData.GetReason())
	tokenLog.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(tokenLog)
	return
}

func init() {
	gameevent.AddEventListener(jieyieventtypes.JieYiEventTypeTokenTypeChangeLog, event.EventListenerFunc(playerJieYiTokenLog))
}
