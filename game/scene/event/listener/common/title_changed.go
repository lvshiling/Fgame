package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//称号改变
func titleChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	if pl.GetScene() == nil {
		return
	}
	if pl.GetScene().MapTemplate().IsTower() {
		return
	}
	scenePlayerTitleChanged := pbutil.BuildScenePlayerTitleChanged(pl)
	scenelogic.BroadcastNeighborIncludeSelf(pl, scenePlayerTitleChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowTitleChanged, event.EventListenerFunc(titleChanged))
}
