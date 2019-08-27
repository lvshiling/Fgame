package listener

import (
	"fgame/fgame/core/event"
	tulonglogic "fgame/fgame/cross/tulong/logic"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
)

//采集龙蛋 被攻击打断
func playerAttacked(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	tulonglogic.CollectEggInterrupt(pl)
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectAttacked, event.EventListenerFunc(playerAttacked))
}
