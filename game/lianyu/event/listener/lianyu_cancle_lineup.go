package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lianyueventtypes "fgame/fgame/game/lianyu/event/types"
	lianyulogic "fgame/fgame/game/lianyu/logic"
)

//玩家取消排队
func lianYuCancleLineUp(target event.EventTarget, data event.EventData) (err error) {
	lineList := target.([]int64)
	pos := data.(int32)
	lianyulogic.BroadLianYuLineUpChanged(pos, lineList)
	return
}

func init() {
	gameevent.AddEventListener(lianyueventtypes.EventTypeLianYuCancleLineUp, event.EventListenerFunc(lianYuCancleLineUp))
}
