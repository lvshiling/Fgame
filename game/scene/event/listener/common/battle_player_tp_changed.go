package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//玩家属性变化
func battlePlayerTPChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	scScenePlayerTPChanged := pbutil.BuildSCScenePlayerTPChanged(pl)
	pl.SendMsg(scScenePlayerTPChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerTPChanged, event.EventListenerFunc(battlePlayerTPChanged))
}
