package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/lineup/lineup"
	lineuplogic "fgame/fgame/cross/lineup/logic"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
)

//场景结束
func sceneFinish(target event.EventTarget, data event.EventData) (err error) {
	s, ok := target.(scene.Scene)
	if !ok {
		return
	}

	crossType, ok := s.MapTemplate().GetMapType().ToCrossType()
	if !ok {
		return
	}

	lineList := lineup.GetLineupService().GetAllLineUpList(crossType, s.Id())
	lineuplogic.BroadLineUpFinishToCancel(int32(crossType), lineList)

	lineup.GetLineupService().ClearAllLineupList(crossType, s.Id())
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeSceneFinish, event.EventListenerFunc(sceneFinish))
}
