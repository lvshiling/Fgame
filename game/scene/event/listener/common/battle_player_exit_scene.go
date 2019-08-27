package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

//玩家退出场景
func battlePlayerExitScene(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	// log.Info("玩家退出场景,同步视野")
	//同步一下视野 防止下个场景出现看不到玩家
	scenelogic.PlayerSyncNeighbors(pl)
	//同步加载过的玩家
	scenelogic.PlayerSyncLoadedPlayers(pl)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerExitScene, event.EventListenerFunc(battlePlayerExitScene))
}
