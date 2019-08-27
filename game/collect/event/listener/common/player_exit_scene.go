package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	collectlogic "fgame/fgame/game/collect/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//采集 移动打断
func battlePlayerExitScene(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	n, flag := pl.HasCollect()
	if !flag {
		return
	}
	collectlogic.CollectInterrupt(pl, n)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerExitScene,event.EventListenerFunc(battlePlayerExitScene))
}
