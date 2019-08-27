package listener

import (
	"fgame/fgame/core/event"
	pbuitl "fgame/fgame/cross/collect/pbutil"
	collecteventtypes "fgame/fgame/game/collect/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//采集 采集完成
func playerCollectFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}

	n, ok := data.(scene.NPC)
	if !ok {
		return
	}

	npcId := n.GetId()
	biologyId := int32(n.GetBiologyTemplate().TemplateId())

	scSceneCollectFinish := pbuitl.BuildSCSceneCollectFinish(npcId, biologyId, nil)
	pl.SendMsg(scSceneCollectFinish)

	isCollectFinish := pbuitl.BuildISCollectFinish(biologyId)
	pl.SendMsg(isCollectFinish)
	return
}

func init() {
	gameevent.AddEventListener(collecteventtypes.EventTypeCollectFinish, event.EventListenerFunc(playerCollectFinish))
}
