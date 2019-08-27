package common

import (
	"fgame/fgame/core/event"
	collecteventtypes "fgame/fgame/game/collect/event/types"
	collectnpc "fgame/fgame/game/collect/npc"
	pbuitl "fgame/fgame/game/collect/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//采集密藏采集物完成
func playerCollectMiZangFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}

	n, ok := data.(scene.NPC)
	if !ok {
		return
	}
	miZangNpc, ok := n.(*collectnpc.CollecMiZangNPC)
	if !ok {
		return
	}
	npcId := n.GetId()
	biologyId := int32(n.GetBiologyTemplate().TemplateId())
	parentId := miZangNpc.GetParentId()
	scMsg := pbuitl.BuildSCSceneCollectMiZangFinish(npcId, biologyId, parentId)
	pl.SendMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(collecteventtypes.EventTypeCollectMiZangFinish, event.EventListenerFunc(playerCollectMiZangFinish))
}
