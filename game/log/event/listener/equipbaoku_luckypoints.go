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

//玩家宝库幸运值变化日志
func playerEquipBaoKuLuckyPointsLog(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*equipbaokueventtypes.PlayerEquipBaoKuLuckyPointsLogEventData)
	if !ok {
		return
	}

	logObj := &logmodel.PlayerEquipBaoKuLuckyPoints{}
	logObj.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logObj.BefNum = eventData.GetBeforeNum()
	logObj.CurNum = eventData.GetCurNum()
	logObj.WithItems = eventData.GetWithItems()
	logObj.Reason = int32(eventData.GetReason())
	logObj.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logObj)
	return
}

func init() {
	gameevent.AddEventListener(equipbaokueventtypes.EventTypeEquipBaoKuLuckyPointsLog, event.EventListenerFunc(playerEquipBaoKuLuckyPointsLog))
}
