package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	godsiegelogic "fgame/fgame/game/godsiege/logic"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
)

//玩家下线
func playerLogout(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	flag := pl.IsGodSiegeLineUp()
	if !flag {
		return
	}
	_, godType := pl.GetGodSiegeLineUp()
	godsiegelogic.GodSiegeCancleLineUpSend(pl, int32(godType))
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerLogout, event.EventListenerFunc(playerLogout))
}
