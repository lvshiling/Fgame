package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	outlandbosseventtypes "fgame/fgame/game/outlandboss/event/types"
	"fgame/fgame/game/player"
)

//浊气值变更
func zhuoQiChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	zhuoQiNum, ok := data.(int32)
	if !ok {
		return
	}

	pl.SynZhuoQiNum(zhuoQiNum)
	return
}

func init() {
	gameevent.AddEventListener(outlandbosseventtypes.EventTypeZhuoQiChanged, event.EventListenerFunc(zhuoQiChanged))
}
