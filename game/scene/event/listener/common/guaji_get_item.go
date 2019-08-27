package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

//挂机获取
func guaJiGetItem(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	s := pl.GetScene()
	if s == nil {
		return
	}
	item := data.(scene.DropItem)
	scenelogic.GetDropItem(pl, item)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerGuaJiGetItem, event.EventListenerFunc(guaJiGetItem))
}
