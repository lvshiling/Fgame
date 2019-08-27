package listener

import (
	"fgame/fgame/core/event"
	dingshieventtypes "fgame/fgame/game/dingshi/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

//进入场景
func dingShiInit(target event.EventTarget, data event.EventData) (err error) {
	s, ok := target.(scene.Scene)
	if !ok {
		return
	}
	npc, ok := data.(scene.NPC)
	if !ok {
		return
	}

	scenelogic.NPCEnterScene(npc, s, npc.GetPosition())
	return
}

func init() {
	gameevent.AddEventListener(dingshieventtypes.EventTypeDingShiInit, event.EventListenerFunc(dingShiInit))
}
