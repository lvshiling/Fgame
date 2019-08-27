package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	xiantieventtypes "fgame/fgame/game/xianti/event/types"
	playerxianti "fgame/fgame/game/xianti/player"
	logmodel "fgame/fgame/logserver/model"
)

//玩家仙体进阶日志
func playerXianTiAdvancedLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*xiantieventtypes.PlayerXianTiAdvancedLogEventData)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	xianTiInfo := manager.GetXianTiInfo()
	logXianTiAdvanced := &logmodel.PlayerXianTi{}
	logXianTiAdvanced.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logXianTiAdvanced.BeforeAdvancedNum = eventData.GetBeforeAdvancedNum()
	logXianTiAdvanced.CurAdvancedNum = int32(xianTiInfo.AdvanceId)
	logXianTiAdvanced.AdvancedNum = eventData.GetAdvancedNum()
	logXianTiAdvanced.Reason = int32(eventData.GetReason())
	logXianTiAdvanced.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logXianTiAdvanced)
	return
}

func init() {
	gameevent.AddEventListener(xiantieventtypes.EventTypeXianTiAdvancedLog, event.EventListenerFunc(playerXianTiAdvancedLog))
}
