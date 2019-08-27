package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
)

//玩家下线
func playerLogout(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl.GetAllianceId() == 0 {
		return
	}
	alliance.GetAllianceService().OfflineMember(pl.GetId())
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerLogout, event.EventListenerFunc(playerLogout))
}
