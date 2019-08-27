package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	bufflogic "fgame/fgame/game/buff/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//攻击
func battlePlayerCollect(target event.EventTarget, data event.EventData) (err error) {

	bo := target.(scene.Player)

	bufflogic.RemoveBuffByAction(bo, scenetypes.BuffRemoveTypeCollect)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerCollect, event.EventListenerFunc(battlePlayerCollect))
}
