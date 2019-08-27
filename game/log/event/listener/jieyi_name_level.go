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

//玩家结义威名等级改变日志
func playerJieYiNameLevelLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*jieyieventtypes.PlayerJieYiNameLevelChangeLogEventData)
	if !ok {
		return
	}

	nameLevelLog := &logmodel.PlayerJieYiNameLevel{}
	nameLevelLog.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	nameLevelLog.BeforeNameLevel = eventData.GetBeformLevel()
	nameLevelLog.CurNameLevel = eventData.GetCurLevel()
	nameLevelLog.Reason = int32(eventData.GetReason())
	nameLevelLog.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(nameLevelLog)
	return
}

func init() {
	gameevent.AddEventListener(jieyieventtypes.JieYiEventTypeNameLevelChangeLog, event.EventListenerFunc(playerJieYiNameLevelLog))
}
