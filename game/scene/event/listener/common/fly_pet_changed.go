package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//宠物变化
func flyPetChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	if pl.GetScene() == nil {
		return
	}

	scenePlayerPetChanged := pbutil.BuildScenePlayerFlyPetChanged(pl)
	scenelogic.BroadcastNeighborIncludeSelf(pl, scenePlayerPetChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowFlyPetChanged, event.EventListenerFunc(flyPetChanged))
}
