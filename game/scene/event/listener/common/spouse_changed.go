package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//配偶改变
func spouseChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	if pl.GetScene() == nil {
		return
	}

	scenePlayerSpouseChanged := pbutil.BuildScenePlayerSpouseChanged(pl)
	scenelogic.BroadcastNeighborIncludeSelf(pl, scenePlayerSpouseChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowSpouseChanged, event.EventListenerFunc(spouseChanged))
}
