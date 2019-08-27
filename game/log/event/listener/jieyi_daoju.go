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

//玩家结义道具改变日志
func playerJieYiDaoJuLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*jieyieventtypes.PlayerJieYiDaoJuChangeLogEventData)
	if !ok {
		return
	}

	daoJuLog := &logmodel.PlayerJieYiDaoJu{}
	daoJuLog.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	daoJuLog.BeforeDaoJu = eventData.GetBeformType().String()
	daoJuLog.CurDaoJu = eventData.GetCurType().String()
	daoJuLog.Reason = int32(eventData.GetReason())
	daoJuLog.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(daoJuLog)
	return
}

func init() {
	gameevent.AddEventListener(jieyieventtypes.JieYiEventTypeDaoJuTypeChangeLog, event.EventListenerFunc(playerJieYiDaoJuLog))
}
