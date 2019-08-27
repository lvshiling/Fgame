package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/jieyi/jieyi"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
)

//玩家登出
func playerLogout(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)

	jieyi.GetJieYiService().PlayerLogout(pl.GetId())
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerLogout, event.EventListenerFunc(playerLogout))
}
