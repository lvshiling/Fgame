package listener

import (
	"fgame/fgame/core/event"
	chuangshiscene "fgame/fgame/cross/chuangshi/scene"
	collecteventtypes "fgame/fgame/game/collect/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//采集 采集完成
func playerCollectFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}

	collectNpc, ok := data.(scene.NPC)
	if !ok {
		return
	}

	s := pl.GetScene()
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeChuangShiZhiZhanFuShu {
		return
	}

	// TODO xzk27 移到场景里面

	chuangshiSd, ok := s.SceneDelegate().(chuangshiscene.FuShuSceneData)
	if !ok {
		return
	}

	if !chuangshiSd.YuXiNpcCollectFinish(collectNpc, pl.GetCamp()) {
		return
	}

	// scMsg := pbutil.BuildSCAllianceSceneOccupyFinish(chuangshiId, name, huFu)
	// s.BroadcastMsg(scMsg)

	return
}

func init() {
	gameevent.AddEventListener(collecteventtypes.EventTypeCollectFinish, event.EventListenerFunc(playerCollectFinish))
}
