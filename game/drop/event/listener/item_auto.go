package listener

import (
	"fgame/fgame/core/event"
	dropeventtypes "fgame/fgame/game/drop/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

//物品自动获取
func dropItemAuto(target event.EventTarget, data event.EventData) (err error) {
	dropItem := target.(scene.DropItem)
	_, err = scenelogic.AutoGetDropItem(dropItem)
	if err != nil {
		return
	}
	return
}

func init() {
	gameevent.AddEventListener(dropeventtypes.EventTypeDropItemAuto, event.EventListenerFunc(dropItemAuto))
}
