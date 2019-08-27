package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	friendeventtypes "fgame/fgame/game/friend/event/types"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	logmodel "fgame/fgame/logserver/model"
)

// 赠送好友装备日志
func friendGiveLog(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*friendeventtypes.FriendGiveEventData)
	if !ok {
		return
	}

	friendLog := &logmodel.FriendGive{}
	friendLog.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	friendLog.Reason = int32(eventData.GetReason())
	friendLog.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(friendLog)
	return
}

func init() {
	gameevent.AddEventListener(friendeventtypes.EventTypeFriendGiveLog, event.EventListenerFunc(friendGiveLog))
}
