package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	bufflogic "fgame/fgame/game/buff/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//pk打断
func playerPKStateSwitch(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(scene.Player)
	if !ok {
		return
	}

	bufflogic.RemoveBuffByAction(p, scenetypes.BuffRemoveTypeSwitchPK)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerPkStateSwitch, event.EventListenerFunc(playerPKStateSwitch))
}
