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

//银两变化
func playerSilverChanged(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*propertyeventtypes.PlayerSilverChangedLogEventData)
	if !ok {
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	curSilver := propertyManager.GetSilver()
	beforeSilver := eventData.GetBeforeSilver()

	logSilverChanged := &logmodel.PlayerSilverChanged{}
	logSilverChanged.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logSilverChanged.ChangedNum = eventData.GetChangedNum()
	logSilverChanged.BeforeSilver = beforeSilver
	logSilverChanged.CurSilver = curSilver
	logSilverChanged.Reason = eventData.GetReason().Reason()
	logSilverChanged.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logSilverChanged)
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerSilverChangedLog, event.EventListenerFunc(playerSilverChanged))
}
