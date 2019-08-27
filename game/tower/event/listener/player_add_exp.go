package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	propertyeventypes "fgame/fgame/game/property/event/types"
	"fgame/fgame/game/scene/scene"
)

//TODO 修改为杀怪获得的经验事件
//打宝塔统计经验
func playerExpAdd(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	s := pl.GetScene()
	if s == nil {
		return
	}
	expAdd := data.(int64)
	if !s.MapTemplate().IsTower() {
		return
	}

	pl.CountTowerExp(expAdd)
	return
}

func init() {
	gameevent.AddEventListener(propertyeventypes.EventTypePlayerExpAdd, event.EventListenerFunc(playerExpAdd))
}
