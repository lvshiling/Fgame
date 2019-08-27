package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

//同步玩家
func playerSyncNeighbor(target event.EventTarget, data event.EventData) (err error) {
	p := target.(scene.Player)
	if p.GetScene() == nil {
		return
	}
	scenelogic.PlayerSyncNeighbors(p)
	return nil
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypePlayerSyncNeighbor, event.EventListenerFunc(playerSyncNeighbor))
}
