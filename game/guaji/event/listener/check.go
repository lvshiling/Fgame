package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	guajieventtypes "fgame/fgame/game/guaji/event/types"
	"fgame/fgame/game/guaji/guaji"
	"fgame/fgame/game/player"
)

//检查挂机提升
func guaJiCheck(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	for _, h := range guaji.GetGuaJiCheckHandlerMap() {
		h.Check(pl)
	}
	return
}

func init() {
	gameevent.AddEventListener(guajieventtypes.GuaJiEventTypeGuaJiCheck, event.EventListenerFunc(guaJiCheck))
}
