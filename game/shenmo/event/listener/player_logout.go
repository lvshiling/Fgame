package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	shenmologic "fgame/fgame/game/shenmo/logic"
)

//玩家下线
func playerLogout(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	flag := pl.IsShenMoLineUp()
	if !flag {
		return
	}
	shenmologic.ShenMoCancleLineUpSend(pl)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerLogout, event.EventListenerFunc(playerLogout))
}
