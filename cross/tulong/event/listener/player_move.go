package listener

import (
	"fgame/fgame/core/event"
	tulonglogic "fgame/fgame/cross/tulong/logic"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
)

//采集龙蛋 移动打断
func battleObjectMove(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	tulonglogic.CollectEggInterrupt(pl)
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectMove, event.EventListenerFunc(battleObjectMove))
}
