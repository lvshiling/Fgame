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
func realmChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	s := pl.GetScene()
	if s == nil {
		return
	}

	scenePlayerRealmChanged := pbutil.BuildScenePlayerRealmChanged(pl)
	scenelogic.BroadcastNeighborIncludeSelf(pl, scenePlayerRealmChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowRealmChanged, event.EventListenerFunc(realmChanged))
}
