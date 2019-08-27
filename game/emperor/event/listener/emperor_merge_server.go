package listener

import (
	"fgame/fgame/core/event"
	emperoreventtypes "fgame/fgame/game/emperor/event/types"
	emperorlogic "fgame/fgame/game/emperor/logic"
	gameevent "fgame/fgame/game/event"
)

//龙椅合服
func emperorMergeServer(target event.EventTarget, data event.EventData) (err error) {
	playerId, ok := target.(int64)
	if !ok {
		return
	}
	itemMap, ok := data.(map[int32]int32)
	if !ok {
		return
	}
	emperorlogic.MergeServeGiveBack(playerId, itemMap)
	return
}

func init() {
	gameevent.AddEventListener(emperoreventtypes.EmperorMergeServer, event.EventListenerFunc(emperorMergeServer))
}
