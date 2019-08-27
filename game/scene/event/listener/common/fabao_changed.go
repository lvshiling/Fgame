package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//身法改变
func faBaoChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	if pl.GetScene() == nil {
		return
	}

	scenePlayerFaBaoChanged := pbutil.BuildScenePlayerFaBaoChanged(pl)
	scenelogic.BroadcastNeighborIncludeSelf(pl, scenePlayerFaBaoChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowFaBaoChanged, event.EventListenerFunc(faBaoChanged))
}
