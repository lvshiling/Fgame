package listener

import (
	"fgame/fgame/core/event"
	collecteventtypes "fgame/fgame/game/collect/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//玩家采集完成
func playerCollectFinish(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	s := pl.GetScene()
	//采集完成场景 刚好结束
	if s == nil {
		return
	}
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeCrossDenseWat {
		return
	}
	num := pl.GetDenseWatNum()
	pl.SetDenseWatNum(num + 1)

	return
}

func init() {
	gameevent.AddEventListener(collecteventtypes.EventTypeCollectFinish, event.EventListenerFunc(playerCollectFinish))
}
