package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/log/log"
	gamelog "fgame/fgame/game/log/log"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	logmodel "fgame/fgame/logserver/model"
)

//在线人数同步
func onlineNumSync(target event.EventTarget, data event.EventData) (err error) {

	systemOnline := &logmodel.SystemOnline{}
	systemOnline.SystemLogMsg = gamelog.SystemLogMsg()
	systemOnline.OnlineNum = player.GetOnlinePlayerManager().Count()

	log.GetLogService().SendLog(systemOnline)

	systemNeiGuaOnline := &logmodel.SystemNeiguaOnline{}
	systemNeiGuaOnline.SystemLogMsg = gamelog.SystemLogMsg()
	systemNeiGuaOnline.OnlineNum = player.GetOnlinePlayerManager().CountNeigua()

	log.GetLogService().SendLog(systemNeiGuaOnline)

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypeOnlineNumSync, event.EventListenerFunc(onlineNumSync))
}
