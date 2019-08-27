package listener

import (
	"fgame/fgame/core/event"
	wingeventtypes "fgame/fgame/game/wing/event/types"
	playerwing "fgame/fgame/game/wing/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家战翼进阶日志
func playerWingAdvancedLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*wingeventtypes.PlayerWingAdvancedLogEventData)
	if !ok {
		return
	}

	wingManager := pl.GetPlayerDataManager(playertypes.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	wingInfo := wingManager.GetWingInfo()

	logWingAdvanced := &logmodel.PlayerWing{}
	logWingAdvanced.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logWingAdvanced.BeforeAdvancedNum = eventData.GetBeforeAdvancedNum()
	logWingAdvanced.CurAdvancedNum = int32(wingInfo.AdvanceId)
	logWingAdvanced.AdvancedNum = eventData.GetAdvancedNum()
	logWingAdvanced.Reason = int32(eventData.GetReason())
	logWingAdvanced.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logWingAdvanced)
	return
}

func init() {
	gameevent.AddEventListener(wingeventtypes.EventTypeWingAdvancedLog, event.EventListenerFunc(playerWingAdvancedLog))
}
