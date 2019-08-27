package common

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lianyueventtypes "fgame/fgame/game/lianyu/event/types"
	"fgame/fgame/game/lianyu/pbutil"
	"fgame/fgame/game/scene/scene"
)

//无间炼狱 玩家杀气变化
func lianYuPlayerShaQiChanged(target event.EventTarget, data event.EventData) (err error) {
	spl, ok := target.(scene.Player)
	if !ok {
		return
	}
	shaQi, ok := data.(int32)
	if !ok {
		return
	}

	scMsg := pbutil.BuildSCLianYuShaQiChanged(shaQi)
	spl.SendMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(lianyueventtypes.EventTypeLianYuPlayerShaQiChanged, event.EventListenerFunc(lianYuPlayerShaQiChanged))
}
