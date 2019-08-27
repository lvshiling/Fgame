package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	playerproperty "fgame/fgame/game/property/player"
	logmodel "fgame/fgame/logserver/model"
)

//新元宝变化
func playerNewGoldChanged(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*propertyeventtypes.PlayerNewGoldChangedLogEventData)
	if !ok {
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	curGold := propertyManager.GetGold()
	beforeGold := eventData.GetBeforeGold()

	logGoldChanged := &logmodel.PlayerNewGoldChanged{}
	logGoldChanged.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logGoldChanged.ChangedNum = eventData.GetChangedNum()
	logGoldChanged.BeforeGold = beforeGold
	logGoldChanged.CurGold = curGold
	logGoldChanged.Reason = eventData.GetReason().Reason()
	logGoldChanged.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logGoldChanged)
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerNewGoldChangedLog, event.EventListenerFunc(playerNewGoldChanged))
}
