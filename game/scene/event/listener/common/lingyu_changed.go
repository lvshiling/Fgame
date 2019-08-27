package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//领域改变
func lingyuChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	if pl.GetScene() == nil {
		return
	}

	scenePlayerLingyuChanged := pbutil.BuildScenePlayerLingyuChanged(pl)
	scenelogic.BroadcastNeighborIncludeSelf(pl, scenePlayerLingyuChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowLingYuChanged, event.EventListenerFunc(lingyuChanged))
}
