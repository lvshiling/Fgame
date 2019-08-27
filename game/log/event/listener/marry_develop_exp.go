package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	"fgame/fgame/game/player"
	logmodel "fgame/fgame/logserver/model"
)

//表白经验升级等级日志
func playerMarryDevelopExpByItemLog(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*marryeventtypes.PlayerDevelopExpLogEventData)
	if !ok {
		return
	}

	playerMarryDevelopExp := &logmodel.PlayerMarryDevelopExp{}
	playerMarryDevelopExp.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	playerMarryDevelopExp.BeforeExp = eventData.GetBeforeDevelopExp()
	playerMarryDevelopExp.CurExp = eventData.GetCurDevelopExp()
	playerMarryDevelopExp.ChangedExp = eventData.GetCurDevelopExp() - eventData.GetBeforeDevelopExp()
	playerMarryDevelopExp.Reason = int32(eventData.GetReason())
	playerMarryDevelopExp.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(playerMarryDevelopExp)
	return
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypePlayerMarryDevelopExpLog, event.EventListenerFunc(playerMarryDevelopExpByItemLog))
}
