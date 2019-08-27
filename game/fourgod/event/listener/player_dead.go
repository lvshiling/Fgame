package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	fourgodlogic "fgame/fgame/game/fourgod/logic"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
)

//四神遗迹宝箱采集 死亡打断
func playerDead(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	fourgodlogic.CollectBoxInterrupt(pl)
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(playerDead))
}
