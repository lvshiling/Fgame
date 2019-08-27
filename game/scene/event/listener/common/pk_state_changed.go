package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	pkpbutil "fgame/fgame/game/pk/pbutil"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//pk状态改变
func pkStateSwitch(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	if pl.GetScene() == nil {
		return
	}
	pkState := pl.GetPkState()
	camp := pl.GetPkCamp()
	pkStateSwitch := pbutil.BuildScenePlayerPkStateSwitched(pl, pkState, camp)

	scenelogic.BroadcastNeighborIncludeSelf(pl, pkStateSwitch)
	scPkStateSwitch := pkpbutil.BuildSCPkStateSwitch(pkState)
	pl.SendMsg(scPkStateSwitch)

	for _, obj := range pl.GetNeighbors() {
		n, ok := obj.(scene.NPC)
		if !ok {
			continue
		}
		isEnemy := pl.IsEnemy(n)
		scMonsterCampChanged := pbutil.BuildSCMonsterCampChanged(n, isEnemy)
		pl.SendMsg(scMonsterCampChanged)
	}

	return nil
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerPkStateSwitch, event.EventListenerFunc(pkStateSwitch))
}
