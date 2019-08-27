package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	exceptioneventtypes "fgame/fgame/game/exception/event/types"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	logmodel "fgame/fgame/logserver/model"
)

//异常错误
func exception(target event.EventTarget, data event.EventData) (err error) {
	content, ok := data.(string)
	if !ok {
		return
	}
	exc := &logmodel.Exception{}
	exc.SystemLogMsg = gamelog.SystemLogMsg()
	exc.Content = content

	log.GetLogService().SendLog(exc)
	return
}

func init() {
	gameevent.AddEventListener(exceptioneventtypes.ExceptionEventTypeException, event.EventListenerFunc(exception))
}
