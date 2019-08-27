package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//玩家战力变化
func playerForceChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	if p.GetAllianceId() == 0 {
		return
	}
	alliance.GetAllianceService().UpdateForce(p.GetId(), p.GetForce())
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerForceChanged, event.EventListenerFunc(playerForceChanged))
}
