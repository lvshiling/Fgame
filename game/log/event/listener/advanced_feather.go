package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	wingeventtypes "fgame/fgame/game/wing/event/types"
	playerwing "fgame/fgame/game/wing/player"
	logmodel "fgame/fgame/logserver/model"
)

//玩家仙羽进阶日志
func playerFeatherAdvancedLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*wingeventtypes.PlayerFeatherAdvancedLogEventData)
	if !ok {
		return
	}

	wingManager := pl.GetPlayerDataManager(playertypes.PlayerWingDataManagerType).(*playerwing.PlayerWingDataManager)
	wingInfo := wingManager.GetWingInfo()

	logFeatherAdvanced := &logmodel.PlayerFeather{}
	logFeatherAdvanced.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logFeatherAdvanced.BeforeAdvancedNum = eventData.GetBeforeAdvancedNum()
	logFeatherAdvanced.CurAdvancedNum = int32(wingInfo.FeatherId)
	logFeatherAdvanced.AdvancedNum = eventData.GetAdvancedNum()
	logFeatherAdvanced.Reason = int32(eventData.GetReason())
	logFeatherAdvanced.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logFeatherAdvanced)
	return
}

func init() {
	gameevent.AddEventListener(wingeventtypes.EventTypeFeatherAdvancedLog, event.EventListenerFunc(playerFeatherAdvancedLog))
}
