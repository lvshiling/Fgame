package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/alliance/alliance"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	propertyeventtypes "fgame/fgame/game/property/event/types"
)

//玩家转生变化
func playerZhuanShengChanged(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	if p.GetAllianceId() == 0 {
		return
	}
	zhuanSheng := data.(int32)
	alliance.GetAllianceService().UpdateZhuanSheng(p.GetId(), zhuanSheng)
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerZhuanShengChanged, event.EventListenerFunc(playerZhuanShengChanged))
}
