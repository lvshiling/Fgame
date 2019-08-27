package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//复活刷新
func playerVipChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(scene.Player)
	if !ok {
		return
	}
	if p.GetScene() == nil {
		return
	}

	scenePlayerVipChanged := pbutil.BuildScenePlayerVipChanged(p)
	scenelogic.BroadcastNeighborIncludeSelf(p, scenePlayerVipChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerVipChanged, event.EventListenerFunc(playerVipChanged))
}
