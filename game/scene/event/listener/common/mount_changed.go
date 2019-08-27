package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	battlelogic "fgame/fgame/game/battle/logic"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//战翼改变
func mountChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	if pl.GetScene() == nil {
		return
	}

	scenePlayerMountChanged := pbutil.BuildScenePlayerMountChanged(pl)
	scenelogic.BroadcastNeighborIncludeSelf(pl, scenePlayerMountChanged)
	battlelogic.UpdateMountBattleProperty(pl)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowMountChanged, event.EventListenerFunc(mountChanged))
}
