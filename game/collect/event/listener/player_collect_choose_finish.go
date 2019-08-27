package listener

import (
	"fgame/fgame/core/event"
	collecteventtypes "fgame/fgame/game/collect/event/types"
	collectlogic "fgame/fgame/game/collect/logic"
	pbuitl "fgame/fgame/game/collect/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
)

//采集 采集物选择完成
func playerCollectChooseFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}

	finishData, ok := data.(*collecteventtypes.CollectChooseFinishEventData)
	if !ok {
		return
	}

	n := finishData.GetCollectNpc()
	typ := finishData.GetChooseFinishType()

	npcId := n.GetId()
	biologyId := int32(n.GetBiologyTemplate().TemplateId())
	tpl := pl.(player.Player)

	dropItemList := collectlogic.CollectChooseDropToInventory(tpl, biologyId, typ)

	scSceneCollectFinish := pbuitl.BuildSCSceneCollectFinish(npcId, biologyId, dropItemList)
	tpl.SendMsg(scSceneCollectFinish)
	return
}

func init() {
	gameevent.AddEventListener(collecteventtypes.EventTypeCollectChooseFinish, event.EventListenerFunc(playerCollectChooseFinish))
}
