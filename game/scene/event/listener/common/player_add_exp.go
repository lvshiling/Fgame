package common

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	propertyeventypes "fgame/fgame/game/property/event/types"
	"fgame/fgame/game/scene/scene"
)

//TODO 修改为杀怪获得的经验事件
//玩家增加经验
func playerExpAdd(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	if pl.GetScene() == nil {
		return
	}
	expAdd := data.(int64)
	pl.GetScene().OnPlayerGetExp(pl, expAdd)
	return nil
}

func init() {
	gameevent.AddEventListener(propertyeventypes.EventTypePlayerExpAdd, event.EventListenerFunc(playerExpAdd))
}
