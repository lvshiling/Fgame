package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//钥匙改变
func fourGodKeyChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	s := pl.GetScene()
	if s == nil {
		return
	}

	scenePlayerFourGodKeyChanged := pbutil.BuildScenePlayerFourGodKeyChanged(pl)
	scenelogic.BroadcastNeighborIncludeSelf(pl, scenePlayerFourGodKeyChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowFourGodKeyChanged, event.EventListenerFunc(fourGodKeyChanged))
}
