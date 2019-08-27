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
func shenfaChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	if pl.GetScene() == nil {
		return
	}

	scenePlayerShenfaChanged := pbutil.BuildScenePlayerShenfaChanged(pl)
	scenelogic.BroadcastNeighborIncludeSelf(pl, scenePlayerShenfaChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowShenFaChanged, event.EventListenerFunc(shenfaChanged))
}
