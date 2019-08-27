package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	unrealbosseventtypes "fgame/fgame/game/unrealboss/event/types"
)

//疲劳值变更
func pilaoChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	pilaoNum, ok := data.(int32)
	if !ok {
		return
	}

	pl.SynPilaoNum(pilaoNum)
	return
}

func init() {
	gameevent.AddEventListener(unrealbosseventtypes.EventTypeUnrealPilaoChanged, event.EventListenerFunc(pilaoChanged))
}
