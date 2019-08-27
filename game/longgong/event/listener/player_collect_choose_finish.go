package listener

import (
	"fgame/fgame/core/event"
	collecteventtypes "fgame/fgame/game/collect/event/types"
	gameevent "fgame/fgame/game/event"
	longgonglogic "fgame/fgame/game/longgong/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//采集 采集物选择完成
func playerCollectChooseFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}

	s := pl.GetScene()
	if s == nil {
		return
	}

	mapType := s.MapTemplate().GetMapType()
	if mapType != scenetypes.SceneTypeLongGong {
		return
	}

	finishData, ok := data.(*collecteventtypes.CollectChooseFinishEventData)
	if !ok {
		return
	}

	n := finishData.GetCollectNpc()
	bioType := n.GetBiologyTemplate().GetBiologyScriptType()
	if bioType != scenetypes.BiologyScriptTypePearl {
		return
	}

	sd := s.SceneDelegate()
	longgongSd, ok := sd.(longgonglogic.LongGongSceneData)
	if !ok {
		return
	}

	longgongSd.AddPearlCollectCount(1)
	return
}

func init() {
	gameevent.AddEventListener(collecteventtypes.EventTypeCollectChooseFinish, event.EventListenerFunc(playerCollectChooseFinish))
}
