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

//元宝变化
func playerGoldChanged(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*propertyeventtypes.PlayerGoldChangedLogEventData)
	if !ok {
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	curGold := propertyManager.GetGold()
	curBindGold := propertyManager.GetBindGlod()
	beforeGold := eventData.GetBeforeGold()
	beforeBindGold := eventData.GetBeforeBindGold()

	logGoldChanged := &logmodel.PlayerGoldChanged{}
	logGoldChanged.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logGoldChanged.ChangedNum = eventData.GetChangedNum()
	logGoldChanged.BeforeGold = beforeGold
	logGoldChanged.BeforeBindGold = beforeBindGold
	logGoldChanged.CurGold = curGold
	logGoldChanged.CurBindGold = curBindGold
	logGoldChanged.Reason = eventData.GetReason().Reason()
	logGoldChanged.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logGoldChanged)
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerGoldChangedLog, event.EventListenerFunc(playerGoldChanged))
}
