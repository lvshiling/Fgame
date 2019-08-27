package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//结婚状态改变
func modelChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	s := pl.GetScene()
	if s == nil {
		return
	}

	scenePlayerModelChanged := pbutil.BuildScenePlayerModelChanged(pl)
	scenelogic.BroadcastNeighborIncludeSelf(pl, scenePlayerModelChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowModelChanged, event.EventListenerFunc(modelChanged))
}
