package listener

import (
	"fgame/fgame/core/event"
	lineupeventtypes "fgame/fgame/cross/lineup/event/types"
	lineuplogic "fgame/fgame/cross/lineup/logic"
	gameevent "fgame/fgame/game/event"
)

//玩家取消排队
func lineupCancleLineUp(target event.EventTarget, data event.EventData) (err error) {
	lineList := target.([]int64)
	eventData := data.(*lineupeventtypes.CancleLineUpEventData)

	lineuplogic.BroadLineUpChanged(eventData.GetIndex(), lineList, eventData.GetCrossType())
	return
}

func init() {
	gameevent.AddEventListener(lineupeventtypes.EventTypeLineupCancleLineUp, event.EventListenerFunc(lineupCancleLineUp))
}
