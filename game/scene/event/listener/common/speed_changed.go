package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//速度变化
func speedChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)

	if pl.GetScene() == nil {
		return
	}

	scenePlayerSpeedChanged := pbutil.BuildScenePlayerSpeedChanged(pl)
	scenelogic.BroadcastNeighborIncludeSelf(pl, scenePlayerSpeedChanged)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerSpeedChanged, event.EventListenerFunc(speedChanged))
}
