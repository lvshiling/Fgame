package listener

import (
	"fgame/fgame/core/event"
	equipbaokueventtypes "fgame/fgame/game/equipbaoku/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	logmodel "fgame/fgame/logserver/model"
)

//玩家宝库积分变化日志
func playerEquipBaoKuAttendPointsLog(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*equipbaokueventtypes.PlayerEquipBaoKuAttendPointsLogEventData)
	if !ok {
		return
	}

	logObj := &logmodel.PlayerEquipBaoKuAttendPoints{}
	logObj.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logObj.BefNum = eventData.GetBeforeNum()
	logObj.CurNum = eventData.GetCurNum()
	logObj.ItemId = eventData.GetItemId()
	logObj.ItemCount = eventData.GetItemCount()
	logObj.Reason = int32(eventData.GetReason())
	logObj.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logObj)
	return
}

func init() {
	gameevent.AddEventListener(equipbaokueventtypes.EventTypeEquipBaoKuAttendPointsLog, event.EventListenerFunc(playerEquipBaoKuAttendPointsLog))
}
