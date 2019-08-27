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

//玩家统计
func playerStats(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(string)
	if !ok {
		return
	}

	stats := &logmodel.PlayerStats{}
	stats.PlayerLogMsg = gamelog.PlayerLogMsgFromPlayer(pl)
	stats.Stats = eventData
	log.GetLogService().SendLog(stats)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerStats, event.EventListenerFunc(playerStats))
}
