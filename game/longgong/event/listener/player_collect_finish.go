package listener

import (
	"fgame/fgame/core/event"
	collecteventtypes "fgame/fgame/game/collect/event/types"
	gameevent "fgame/fgame/game/event"
	longgonglogic "fgame/fgame/game/longgong/logic"
	"fgame/fgame/game/longgong/pbutil"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//采集 采集完成
func playerCollectFinishWith(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}

	eventData, ok := data.(*collecteventtypes.CollectFinishWithEventData)
	if !ok {
		return
	}

	n := eventData.GetCollectNpc()
	bioType := n.GetBiologyTemplate().GetBiologyScriptType()
	if bioType != scenetypes.BiologyScriptTypeLongGongTreasure {
		return
	}

	s := pl.GetScene()
	mapType := s.MapTemplate().GetMapType()
	if mapType != scenetypes.SceneTypeLongGong {
		return
	}

	ppl, ok := pl.(player.Player)
	if !ok {
		return
	}

	//增加采集次数
	sd := s.SceneDelegate()
	longgongSd, ok := sd.(longgonglogic.LongGongSceneData)
	if !ok {
		return
	}
	longgongSd.AddPlayerTreasureCollectCount(pl.GetId())
	pCollectCount := longgongSd.GetPlayerTreasureCollectCount(pl.GetId())

	scMsg := pbutil.BuildSCLonggongPlayerValChange(pCollectCount)
	ppl.SendMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(collecteventtypes.EventTypeCollectFinishWith, event.EventListenerFunc(playerCollectFinishWith))
}
