package common

import (
	"fgame/fgame/core/event"
	battleeventypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//玩家仙盟变化
func playerAllianceChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	if pl.GetScene() == nil {
		return
	}
	allianceName := pl.GetAllianceName()
	allianceChanged := pbutil.BuildScenePlayerAllianceChanged(pl, allianceName)
	scenelogic.BroadcastNeighborIncludeSelf(pl, allianceChanged)
	return nil

}

func init() {
	gameevent.AddEventListener(battleeventypes.EventTypeBattlePlayerAllianceChanged, event.EventListenerFunc(playerAllianceChanged))
}
