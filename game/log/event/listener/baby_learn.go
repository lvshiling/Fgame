package listener

import (
	"fgame/fgame/core/event"
	babyeventtypes "fgame/fgame/game/baby/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	logmodel "fgame/fgame/logserver/model"
)

//玩家宝宝读书日志
func playerBabyLearnLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*babyeventtypes.PlayerBabyLevelLogEventData)
	if !ok {
		return
	}

	logBaby := &logmodel.PlayerBabyLearn{}
	logBaby.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logBaby.BeforeBabyLevel = eventData.GetBeforeLevel()
	logBaby.CurBabyLevel = eventData.GetCurLevel()
	dif := logBaby.CurBabyLevel - logBaby.BeforeBabyLevel
	logBaby.ChangedLevel = dif
	logBaby.Reason = int32(eventData.GetReason())
	logBaby.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logBaby)
	return
}

func init() {
	gameevent.AddEventListener(babyeventtypes.EventTypeBabyLearnLog, event.EventListenerFunc(playerBabyLearnLog))
}
