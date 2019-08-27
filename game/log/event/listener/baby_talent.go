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

//玩家宝宝天赋日志
func playerBabyTalentLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*babyeventtypes.PlayerBabyTalentLogEventData)
	if !ok {
		return
	}

	logBaby := &logmodel.PlayerBabyTalent{}
	logBaby.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logBaby.BeforeBabyTalent = eventData.GetBeforeTalent()
	logBaby.CurBabyTalent = eventData.GetCurTalent()
	logBaby.ChangedTalent = eventData.GetChangedTalent()
	logBaby.Reason = int32(eventData.GetReason())
	logBaby.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logBaby)
	return
}

func init() {
	gameevent.AddEventListener(babyeventtypes.EventTypeBabyTalentLog, event.EventListenerFunc(playerBabyTalentLog))
}
