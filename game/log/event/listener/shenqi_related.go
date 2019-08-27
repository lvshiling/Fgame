package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	shenqieventtypes "fgame/fgame/game/shenqi/event/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家神器相关升级日志
func playerShenQiRelatedUpLevelLog(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*shenqieventtypes.PlayerShenQiRelatedUpLevelLogEventData)
	if !ok {
		return
	}

	logMsg := &logmodel.PlayerShenQiRelated{}
	logMsg.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logMsg.BeforeLevel = eventData.GetBefLev()
	logMsg.CurLevel = eventData.GetCurLev()
	logMsg.Reason = int32(eventData.GetReason())
	logMsg.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logMsg)
	return
}

func init() {
	gameevent.AddEventListener(shenqieventtypes.EventTypeShenQiRelatedUpLevelLog, event.EventListenerFunc(playerShenQiRelatedUpLevelLog))
}
