package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	tianmoeventtypes "fgame/fgame/game/tianmo/event/types"
	playertianmo "fgame/fgame/game/tianmo/player"
	logmodel "fgame/fgame/logserver/model"
)

//玩家天魔体进阶日志
func playerTianMoTiAdvancedLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*tianmoeventtypes.PlayerTianMoAdvancedLogEventData)
	if !ok {
		return
	}

	tianmoManager := pl.GetPlayerDataManager(playertypes.PlayerTianMoDataManagerType).(*playertianmo.PlayerTianMoDataManager)
	tianmoInfo := tianmoManager.GetTianMoInfo()

	logTianMoTiAdvanced := &logmodel.PlayerTianMoTi{}
	logTianMoTiAdvanced.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logTianMoTiAdvanced.BeforeAdvancedNum = eventData.GetBeforeAdvancedNum()
	logTianMoTiAdvanced.CurAdvancedNum = int32(tianmoInfo.AdvanceId)
	logTianMoTiAdvanced.AdvancedNum = eventData.GetAdvancedNum()
	logTianMoTiAdvanced.Reason = int32(eventData.GetReason())
	logTianMoTiAdvanced.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logTianMoTiAdvanced)
	return
}

func init() {
	gameevent.AddEventListener(tianmoeventtypes.EventTypeTianMoAdvancedLog, event.EventListenerFunc(playerTianMoTiAdvancedLog))
}
