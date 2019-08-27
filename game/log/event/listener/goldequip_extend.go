package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	goldequipeventtypes "fgame/fgame/game/goldequip/event/types"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	logmodel "fgame/fgame/logserver/model"
)

//玩家金装继承日志
func playerGoldEquipExtendLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*goldequipeventtypes.PlayerGoldEquipExtendLogEventData)
	if !ok {
		return
	}

	logGoldEquipExtend := &logmodel.PlayerGoldEquipExtend{}
	logGoldEquipExtend.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logGoldEquipExtend.BeforeLevel = eventData.GetBeforeLevel()
	logGoldEquipExtend.AfterLevel = eventData.GetAfterLevel()
	logGoldEquipExtend.Reason = int32(eventData.GetReason())
	logGoldEquipExtend.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logGoldEquipExtend)
	return
}

func init() {
	gameevent.AddEventListener(goldequipeventtypes.EventTypeGoldEquipExtendLog, event.EventListenerFunc(playerGoldEquipExtendLog))
}
