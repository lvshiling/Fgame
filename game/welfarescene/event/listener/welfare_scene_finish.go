package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	welfaresceneeventtypes "fgame/fgame/game/welfarescene/event/types"
	"fgame/fgame/game/welfarescene/welfarescene"
)

func welfareSceneFinish(target event.EventTarget, data event.EventData) (err error) {
	groupId, ok := data.(int32)
	if !ok {
		return
	}

	welfarescene.GetWelfareSceneService().WelfareSceneFinish(groupId)
	return
}

func init() {
	gameevent.AddEventListener(welfaresceneeventtypes.EventTypeWelfareSceneFinish, event.EventListenerFunc(welfareSceneFinish))
}
 