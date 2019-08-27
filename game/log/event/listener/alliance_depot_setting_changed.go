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

//仙盟仓库设置变化
func allianceDepotSettingChangedLog(target event.EventTarget, data event.EventData) (err error) {

	al, ok := target.(*alliance.Alliance)
	if !ok {
		return
	}
	eventData, ok := data.(*allianceeventtypes.AllianceDepotSettingChangedLogEventData)
	if !ok {
		return
	}

	logSetting := &logmodel.AllianceDepotSetting{}
	logSetting.AllianceLogMsg = gamelog.AllianceLogMsgFromPlayer(al)
	logSetting.Reason = int32(eventData.GetReason())
	logSetting.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logSetting)
	return
}

func init() {
	gameevent.AddEventListener(allianceeventtypes.EventTypeAllianceDepotSettingChangedLog, event.EventListenerFunc(allianceDepotSettingChangedLog))
}
