package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	towerventtypes "fgame/fgame/game/tower/event/types"
)

//开始打宝塔
func startDaBao(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	pl.StartDaBao()
	return
}

func init() {
	gameevent.AddEventListener(towerventtypes.EventTypeTowerStartDaBao, event.EventListenerFunc(startDaBao))
}
