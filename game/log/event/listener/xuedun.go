package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	xueduneventtypes "fgame/fgame/game/xuedun/event/types"
	logmodel "fgame/fgame/logserver/model"
)

//玩家血盾升阶日志
func playerXueDunUpgradeLog(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*xueduneventtypes.PlayerXueDunUpgradeLogEventData)
	if !ok {
		return
	}

	logXueDunUpgrade := &logmodel.PlayerXueDun{}
	logXueDunUpgrade.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	logXueDunUpgrade.BeforeNumber = eventData.GetBeforeNumber()
	logXueDunUpgrade.BeforeStar = eventData.GetBeforeStar()
	logXueDunUpgrade.CurNumber = eventData.GetNumber()
	logXueDunUpgrade.CurStar = eventData.GetStar()
	logXueDunUpgrade.Reason = int32(eventData.GetReason())
	logXueDunUpgrade.ReasonText = eventData.GetReasonText()
	log.GetLogService().SendLog(logXueDunUpgrade)
	return
}

func init() {
	gameevent.AddEventListener(xueduneventtypes.EventTypeXueDunUpgradeLog, event.EventListenerFunc(playerXueDunUpgradeLog))
}
