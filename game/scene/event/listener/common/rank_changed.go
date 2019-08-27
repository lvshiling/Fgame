package common

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//排行榜更新
func sceneRankChanged(target event.EventTarget, data event.EventData) (err error) {
	s, ok := target.(scene.Scene)
	if !ok {
		return
	}
	r, ok := data.(*scene.SceneRank)
	if !ok {
		return
	}
	for _, p := range s.GetAllPlayers() {
		scSceneRankChanged := pbutil.BuildSCSceneRankChanged(p, r)
		p.SendMsg(scSceneRankChanged)
	}

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeSceneRankChanged, event.EventListenerFunc(sceneRankChanged))
}
