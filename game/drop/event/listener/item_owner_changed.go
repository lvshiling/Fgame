package listener

import (
	"fgame/fgame/core/event"
	dropeventtypes "fgame/fgame/game/drop/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//物品状态改变
func dropItemOwnerChanged(target event.EventTarget, data event.EventData) (err error) {
	dropItem := target.(scene.DropItem)

	for _, obj := range dropItem.GetNeighbors() {
		p, ok := obj.(scene.Player)
		if !ok {
			continue
		}
		canGet := p.IfCanGetDropItem(dropItem)
		scItemOwnerChanged := pbutil.BuildSCItemOwnerChanged(dropItem, canGet)
		p.SendMsg(scItemOwnerChanged)
	}

	return
}

func init() {
	gameevent.AddEventListener(dropeventtypes.EventTypeDropItemOwnerChanged, event.EventListenerFunc(dropItemOwnerChanged))
}
