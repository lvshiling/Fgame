package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	marrylogic "fgame/fgame/game/marry/logic"
	"fgame/fgame/game/marry/marry"
)

//婚礼撒喜糖
func marrySugarTime(target event.EventTarget, data event.EventData) (err error) {
	hunCheNpc := marry.GetMarryService().GetHunCheNpc()
	if hunCheNpc == nil {
		return
	}

	pos := hunCheNpc.GetPosition()
	hunCheObj := hunCheNpc.GetHunCheObject()
	sugarGrade := hunCheObj.SugarGrade
	marrylogic.HunCheDrop(pos, 0, sugarGrade)
	return nil
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypeMarrySugarTime, event.EventListenerFunc(marrySugarTime))
}
