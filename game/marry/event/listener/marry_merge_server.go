package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	marrylogic "fgame/fgame/game/marry/logic"
	"fgame/fgame/game/marry/marry"
)

//结婚合服
func marryMergeServer(target event.EventTarget, data event.EventData) (err error) {
	marryWed, ok := target.(*marry.MarryWedObject)
	if !ok {
		return
	}
	marrylogic.MergeServeGiveBack(marryWed)
	return
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarryMergeServer, event.EventListenerFunc(marryMergeServer))
}
