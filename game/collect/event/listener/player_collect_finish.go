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
	sceneType := pl.GetScene().MapTemplate().GetMapType()

	tpl := pl.(player.Player)
	dropItemList := collectlogic.CollectDropToInventory(tpl, biologyId)
	eventData := collecteventtypes.CreateCollectFinishWithEventData(n, dropItemList)
	gameevent.Emit(collecteventtypes.EventTypeCollectFinishWith, pl, eventData)

	//更新活动采集次数
	activityType, ok := sceneType.ToActivityType()
	if ok {
		pl.UpdateActivityCollect(activityType, biologyId)
	}

	scSceneCollectFinish := pbuitl.BuildSCSceneCollectFinish(npcId, biologyId, dropItemList)
	tpl.SendMsg(scSceneCollectFinish)
	return
}

func init() {
	gameevent.AddEventListener(collecteventtypes.EventTypeCollectFinish, event.EventListenerFunc(playerCollectFinish))
}
