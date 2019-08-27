package common

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//身法改变
func lingTongFaBaoChanged(target event.EventTarget, data event.EventData) (err error) {
	lingTong, ok := target.(scene.LingTong)
	if !ok {
		return
	}
	//不同场景
	if !scenelogic.CheckIfLingTongAndPlayerSameScene(lingTong) {
		return
	}
	scenePlayerFaBaoChanged := pbutil.BuildSceneLingTongFaBaoChanged(lingTong)
	scenelogic.BroadcastNeighborIncludeSelf(lingTong, scenePlayerFaBaoChanged)
	return
}

func init() {
	gameevent.AddEventListener(lingtongeventtypes.EventTypeBattleLingTongShowFaBaoChanged, event.EventListenerFunc(lingTongFaBaoChanged))
}
