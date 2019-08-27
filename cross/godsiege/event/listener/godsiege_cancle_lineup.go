package listener

import (
	"fgame/fgame/core/event"
	godsiegelogic "fgame/fgame/cross/godsiege/logic"
	gameevent "fgame/fgame/game/event"
	godsiegeeventtypes "fgame/fgame/game/godsiege/event/types"
)

//玩家取消排队
func godSiegeCancleLineUp(target event.EventTarget, data event.EventData) (err error) {
	lineList := target.([]int64)
	eventData := data.(*godsiegeeventtypes.GodSiegeCancleLineUpEventData)
	pos := eventData.GetPos()
	godType := eventData.GetGodType()
	godsiegelogic.BroadGodSiegeLineUpChanged(int32(godType), pos, lineList)
	return
}

func init() {
	gameevent.AddEventListener(godsiegeeventtypes.EventTypeGodSiegeCancleLineUp, event.EventListenerFunc(godSiegeCancleLineUp))
}
