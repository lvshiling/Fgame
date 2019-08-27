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
func playerMount(target event.EventTarget, data event.EventData) (err error) {
	p := target.(scene.Player)
	if p.IsMountHidden() {
		return
	}

	bufflogic.RemoveBuffByAction(p, scenetypes.BuffRemoveTypeMount)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowMountHidden, event.EventListenerFunc(playerMount))
}
