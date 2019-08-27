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

//玩家登陆日志
func playerLogin(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	playerLogin := &logmodel.PlayerLogin{}
	playerLogin.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	log.GetLogService().SendLog(playerLogin)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerLogin))
}
