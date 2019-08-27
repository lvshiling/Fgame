package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//神域钥匙改变
func shenYuKeyChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	s := pl.GetScene()
	if s == nil {
		return
	}

	scMsg := pbutil.BuildScenePlayerShenYuChanged(pl)
	scenelogic.BroadcastNeighborIncludeSelf(pl, scMsg)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowShenYuKeyChanged, event.EventListenerFunc(shenYuKeyChanged))
}
