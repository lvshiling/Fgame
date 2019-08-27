package listener

import (
	"fgame/fgame/core/event"
	anqieventtypes "fgame/fgame/game/anqi/event/types"
	playeranqi "fgame/fgame/game/anqi/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家暗器进阶日志
func playerAnqiAdvancedLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*anqieventtypes.PlayerAnqiAdvancedLogEventData)
	if !ok {
		return
	}

	anqiManager := pl.GetPlayerDataManager(playertypes.PlayerAnqiDataManagerType).(*playeranqi.PlayerAnqiDataManager)
	anqiInfo := anqiManager.GetAnqiInfo()

	logAnqiAdvanced := &logmodel.PlayerAnqi{}
	logAnqiAdvanced.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logAnqiAdvanced.BeforeAdvancedNum = eventData.GetBeforeAdvancedNum()
	logAnqiAdvanced.CurAdvancedNum = int32(anqiInfo.AdvanceId)
	logAnqiAdvanced.AdvancedNum = eventData.GetAdvancedNum()
	logAnqiAdvanced.Reason = int32(eventData.GetReason())
	logAnqiAdvanced.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logAnqiAdvanced)
	return
}

func init() {
	gameevent.AddEventListener(anqieventtypes.EventTypeAnqiAdvancedLog, event.EventListenerFunc(playerAnqiAdvancedLog))
}
