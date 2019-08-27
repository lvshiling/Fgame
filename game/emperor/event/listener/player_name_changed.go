package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/emperor/emperor"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
)

//玩家名字变化
func playerNameChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	emperor.GetEmperorService().PlayerNameChanged(pl)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerNameChanged, event.EventListenerFunc(playerNameChanged))
}
