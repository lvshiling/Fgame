package common

import (
	"fgame/fgame/core/event"
	battleeventypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//玩家结义变化
func playerJieYiChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	if pl.GetScene() == nil {
		return
	}
	jieYiName := pl.GetSceneJieYiName()
	jieYiChanged := pbutil.BuildScenePlayerJieYiChanged(pl, jieYiName)
	scenelogic.BroadcastNeighborIncludeSelf(pl, jieYiChanged)
	return nil

}

func init() {
	gameevent.AddEventListener(battleeventypes.EventTypeBattlePlayerJieYiChanged, event.EventListenerFunc(playerJieYiChanged))
}
