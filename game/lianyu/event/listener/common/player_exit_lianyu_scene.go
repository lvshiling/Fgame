package common

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lianyueventtypes "fgame/fgame/game/lianyu/event/types"
	"fgame/fgame/game/lianyu/lianyu"
	lianyuscene "fgame/fgame/game/lianyu/scene"
)

//玩家退出炼狱
func lianYuPlayerExitScene(target event.EventTarget, data event.EventData) (err error) {
	sd, ok := target.(lianyuscene.LianYuSceneData)
	if !ok {
		return
	}
	num := sd.GetScenePlayerNum()
	lianyu.GetLianYuService().RemoveFirstLineUpPlayer(num)
	return
}

func init() {
	gameevent.AddEventListener(lianyueventtypes.EventTypeLianYuPlayerExit, event.EventListenerFunc(lianYuPlayerExitScene))
}
