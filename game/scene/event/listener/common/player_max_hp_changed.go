package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//最大血量变更
func playerMaxHPChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	if pl.GetScene() == nil {
		return
	}
	scenePlayerMaxHPChanged := pbutil.BuildScenePlayerMaxHPChanged(pl)
	scenelogic.BroadcastNeighborIncludeSelf(pl, scenePlayerMaxHPChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerMaxHPChanged, event.EventListenerFunc(playerMaxHPChanged))
}
