package common

import (
	"fgame/fgame/core/event"
	battleeventypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//玩家阵营变化
func playerCampChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	if pl.GetScene() == nil {
		return
	}
	camp := pl.GetCamp()
	guanZhi := pl.GetGuanZhi()
	campChanged := pbutil.BuildScenePlayerCampChanged(pl, camp, guanZhi)
	scenelogic.BroadcastNeighborIncludeSelf(pl, campChanged)
	return nil

}

func init() {
	gameevent.AddEventListener(battleeventypes.EventTypeBattlePlayerCampChanged, event.EventListenerFunc(playerCampChanged))
}
