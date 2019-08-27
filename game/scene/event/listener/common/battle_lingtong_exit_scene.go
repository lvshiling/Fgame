package common

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingtongeventtypes "fgame/fgame/game/lingtong/event/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

//玩家退出场景
func battleLingTongExitScene(target event.EventTarget, data event.EventData) (err error) {
	lingTong := target.(scene.LingTong)

	//同步加载过的玩家
	scenelogic.LingTongSyncLoadedPlayers(lingTong)
	return
}

func init() {
	gameevent.AddEventListener(lingtongeventtypes.EventTypeBattleLingTongExitScene, event.EventListenerFunc(battleLingTongExitScene))
}
