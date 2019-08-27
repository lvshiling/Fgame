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
func pkValueChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)

	curValue := pl.GetPkValue()
	onlineTime := pl.GetPkOnlineTime()
	loginTime := pl.GetPkLoginTime()
	scMsg := pkpbutil.BuildSCPKValueChanged(curValue, onlineTime, loginTime)
	pl.SendMsg(scMsg)
	
	if pl.GetScene() == nil {
		return
	}

	val := pl.GetPkValue()
	pkValueChanged := pbutil.BuildScenePlayerPkValueChanged(pl, val)
	//TODO 同步状态改变
	scenelogic.BroadcastNeighborIncludeSelf(pl, pkValueChanged)

	return nil
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerPkValueChanged, event.EventListenerFunc(pkValueChanged))
}
