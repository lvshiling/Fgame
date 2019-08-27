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

//玩家结义信物等级改变日志
func playerJieYiTokenLevelLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*jieyieventtypes.PlayerJieYiTokenLevelChangeLogEventData)
	if !ok {
		return
	}

	tokenLevelLog := &logmodel.PlayerJieYiTokenLevel{}
	tokenLevelLog.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	tokenLevelLog.BeforeTokenLevel = eventData.GetBeformLevel()
	tokenLevelLog.CurTokenLevel = eventData.GetCurLevel()
	tokenLevelLog.Reason = int32(eventData.GetReason())
	tokenLevelLog.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(tokenLevelLog)
	return
}

func init() {
	gameevent.AddEventListener(jieyieventtypes.JieYiEventTypeTokenLevelChangeLog, event.EventListenerFunc(playerJieYiTokenLevelLog))
}
