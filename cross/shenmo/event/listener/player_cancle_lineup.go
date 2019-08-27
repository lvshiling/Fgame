package listener

import (
	"fgame/fgame/core/event"
	shenmologic "fgame/fgame/cross/shenmo/logic"
	gameevent "fgame/fgame/game/event"
	shenmoeventtypes "fgame/fgame/game/shenmo/event/types"
)

//玩家取消排队
func shenMoCancleLineUp(target event.EventTarget, data event.EventData) (err error) {
	lineList := target.([]int64)
	pos := data.(int32)
	shenmologic.BroadShenMoLineUpChanged(pos, lineList)
	return
}

func init() {
	gameevent.AddEventListener(shenmoeventtypes.EventTypeShenMoCancleLineUp, event.EventListenerFunc(shenMoCancleLineUp))
}
