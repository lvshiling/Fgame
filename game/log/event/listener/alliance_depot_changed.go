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

//仙盟仓库变化变化
func allianceDepotItemChangedLog(target event.EventTarget, data event.EventData) (err error) {

	al, ok := target.(*alliance.Alliance)
	if !ok {
		return
	}
	eventData, ok := data.(*allianceeventtypes.AllianceDepotItemChangedLogEventData)
	if !ok {
		return
	}

	logItemChanged := &logmodel.AllianceDepot{}
	logItemChanged.AllianceLogMsg = gamelog.AllianceLogMsgFromPlayer(al)
	logItemChanged.ItemId = eventData.GetItemId()
	logItemChanged.ChangedNum = eventData.GetChangedNum()
	logItemChanged.Reason = int32(eventData.GetReason())
	logItemChanged.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logItemChanged)
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceDepotItemChangedLog, event.EventListenerFunc(allianceDepotItemChangedLog))
}
