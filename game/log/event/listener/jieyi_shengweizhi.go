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

//玩家结义声威值改变日志
func playerJieYiShengWeiZhiLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*jieyieventtypes.PlayerJieYiShengWeiZhiChangeLogEventData)
	if !ok {
		return
	}

	shengWeiZhiLog := &logmodel.PlayerJieYiShengWeiZhi{}
	shengWeiZhiLog.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	shengWeiZhiLog.BeforeShengWeiZhi = eventData.GetBeformNum()
	shengWeiZhiLog.CurShengWeiZhi = eventData.GetCurNum()
	shengWeiZhiLog.Reason = int32(eventData.GetReason())
	shengWeiZhiLog.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(shengWeiZhiLog)
	return
}

func init() {
	gameevent.AddEventListener(jieyieventtypes.JieYiEventTypeShengWeiZhiChangeLog, event.EventListenerFunc(playerJieYiShengWeiZhiLog))
}
