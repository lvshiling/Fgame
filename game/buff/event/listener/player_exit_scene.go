package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	bufflogic "fgame/fgame/game/buff/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//移动
func playerExitScene(target event.EventTarget, data event.EventData) (err error) {
	p := target.(scene.Player)
	//退出场景
	bufflogic.RemoveBuffByAction(p, scenetypes.BuffRemoveTypeChangeScene)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerExitScene, event.EventListenerFunc(playerExitScene))
}
