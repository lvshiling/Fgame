package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	jieyieventtypes "fgame/fgame/game/jieyi/event/types"
	"fgame/fgame/game/jieyi/jieyi"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	logmodel "fgame/fgame/logserver/model"
)

//离开结义日志
func jieyiLeaveJieYiLog(target event.EventTarget, data event.EventData) (err error) {

	jieYi, ok := target.(*jieyi.JieYi)
	if !ok {
		return
	}
	memberId, ok := data.(int64)
	if !ok {
		return
	}

	jieyiMerge := &logmodel.JieYiLeave{}
	jieyiMerge.JieYiLogMsg = gamelog.JieYiLogMsgFromPlayer(jieYi)
	jieyiMerge.PlayerId = memberId
	log.GetLogService().SendLog(jieyiMerge)
	return
}

func init() {
	gameevent.AddEventListener(jieyieventtypes.JieYiEventTypeLeaveJieYiLog, event.EventListenerFunc(jieyiLeaveJieYiLog))
}
