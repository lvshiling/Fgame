package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	logmodel "fgame/fgame/logserver/model"
)

//仙盟合并日志
func allianceMergeLog(target event.EventTarget, data event.EventData) (err error) {

	al, ok := target.(*alliance.Alliance)
	if !ok {
		return
	}
	inviteAllianceId, ok := data.(int64)
	if !ok {
		return
	}

	allianceMerge := &logmodel.AllianceMerge{}
	allianceMerge.AllianceLogMsg = gamelog.AllianceLogMsgFromPlayer(al)
	allianceMerge.InviteAllianceId = inviteAllianceId
	log.GetLogService().SendLog(allianceMerge)
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceMergeLog, event.EventListenerFunc(allianceMergeLog))
}
