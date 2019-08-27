package listener

import (
	"fgame/fgame/core/event"
	dropeventtypes "fgame/fgame/game/drop/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//物品移除
func dropItemRemove(target event.EventTarget, data event.EventData) (err error) {
	dropItem := target.(scene.DropItem)
	dropItem.GetScene().RemoveSceneObject(dropItem, false)
	return
}

func init() {
	gameevent.AddEventListener(dropeventtypes.EventTypeDropItemRemove, event.EventListenerFunc(dropItemRemove))
}
