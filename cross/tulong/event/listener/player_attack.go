package listener

import (
	"fgame/fgame/core/event"
	tulonglogic "fgame/fgame/cross/tulong/logic"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
)

//龙蛋采集 攻击打断
func playerAttack(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	tulonglogic.CollectEggInterrupt(pl)
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectAttack, event.EventListenerFunc(playerAttack))
}
