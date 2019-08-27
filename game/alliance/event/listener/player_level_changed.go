package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	propertyeventtypes "fgame/fgame/game/property/event/types"
)

//玩家等级变化
func playerLevelChanged(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	if p.GetAllianceId() == 0 {
		return
	}
	level := data.(int32)
	alliance.GetAllianceService().UpdateLevel(p.GetId(), level)
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerLevelChanged, event.EventListenerFunc(playerLevelChanged))
}
