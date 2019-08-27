package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/emperor/emperor"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
)

//玩家性别变化
func playerSexChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	emperor.GetEmperorService().PlayerSexChanged(pl)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerSexChanged, event.EventListenerFunc(playerSexChanged))
}
