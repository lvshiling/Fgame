package pbutil

import (
	logserverlog "fgame/fgame/logserver/log"
	logserverpb "fgame/fgame/logserver/pb"
)

func ConvertFromLogMessage(m *logserverpb.LogMessage) (msg logserverlog.LogMsg, err error) {
	msg, err = logserverlog.Decode(m.GetLogName(), m.GetContent())
	return
}
