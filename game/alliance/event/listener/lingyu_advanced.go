package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	gameevent "fgame/fgame/game/event"
	lingyueventtypes "fgame/fgame/game/lingyu/event/types"
	"fgame/fgame/game/player"
)

//领域进阶
func lingyuAdvanced(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	advanced := data.(int32)
	if p.GetAllianceId() == 0 {
		return
	}
	alliance.GetAllianceService().UpdateLingYu(p.GetId(), advanced)

	return
}

func init() {
	gameevent.AddEventListener(lingyueventtypes.EventTypeLingyuAdvanced, event.EventListenerFunc(lingyuAdvanced))
}
