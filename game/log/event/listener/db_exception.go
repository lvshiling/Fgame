package listener

import (
	"encoding/json"
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	exceptioneventtypes "fgame/fgame/game/exception/event/types"

	gamelog "fgame/fgame/game/log/log"
	logmodel "fgame/fgame/logserver/model"

	log "github.com/Sirupsen/logrus"
)

//db异常错误
func dbException(target event.EventTarget, data event.EventData) (err error) {
	exceptionData, ok := data.(*exceptioneventtypes.DBExceptionEventData)
	if !ok {
		return
	}
	exc := &logmodel.DBException{}
	exc.SystemLogMsg = gamelog.SystemLogMsg()
	exc.TableName = exceptionData.GetTableName()
	content, terr := json.Marshal(exceptionData.GetData())
	if terr != nil {
		log.WithFields(
			log.Fields{
				"tableName": exceptionData.GetTableName(),
				"data":      exceptionData.GetData(),
				"err":       exceptionData.GetError(),
			}).Warn("log:db异常日志")
		return
	}
	exc.Data = string(content)
	exc.Error = exceptionData.GetError()
	gamelog.GetLogService().SendLog(exc)
	return
}

func init() {
	gameevent.AddEventListener(exceptioneventtypes.ExceptionEventTypeDBException, event.EventListenerFunc(dbException))
}
