package listener

import (
	"fgame/fgame/core/event"
	bufflogic "fgame/fgame/game/buff/logic"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//移动
func battleObjectMove(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	bufflogic.RemoveBuffByAction(bo, scenetypes.BuffRemoveTypeWalk)

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectMove, event.EventListenerFunc(battleObjectMove))
}
