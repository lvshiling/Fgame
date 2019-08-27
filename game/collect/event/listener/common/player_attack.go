package common

import (
	"fgame/fgame/core/event"
	collectlogic "fgame/fgame/game/collect/logic"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
)

//采集 攻击打断
func playerAttack(target event.EventTarget, data event.EventData) (err error) {
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
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectAttack, event.EventListenerFunc(playerAttack))
}