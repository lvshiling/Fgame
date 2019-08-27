package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/jieyi/jieyi"
	"fgame/fgame/game/player"
	propertyeventtypes "fgame/fgame/game/property/event/types"
)

//玩家等级变化
func playerLevelChanged(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}

	level := data.(int32)
	jieyi.GetJieYiService().UpdatePlayerLevel(p.GetId(), level)
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerLevelChanged, event.EventListenerFunc(playerLevelChanged))
}
