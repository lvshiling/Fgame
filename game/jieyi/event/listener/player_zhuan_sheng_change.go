package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/jieyi/jieyi"
	"fgame/fgame/game/player"
	propertyeventtypes "fgame/fgame/game/property/event/types"
)

//玩家转生变化
func playerZhuanShengChanged(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)

	zhuanSheng := data.(int32)
	jieyi.GetJieYiService().UpdatePlayerZhuanSheng(p.GetId(), zhuanSheng)
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerZhuanShengChanged, event.EventListenerFunc(playerZhuanShengChanged))
}
