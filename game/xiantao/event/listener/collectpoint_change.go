package listener

import (
	"fgame/fgame/core/event"
	collecteventtypes "fgame/fgame/game/collect/event/types"
	collectnpc "fgame/fgame/game/collect/npc"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/xiantao/pbutil"
)

//采集点变化
func collectPointChange(target event.EventTarget, data event.EventData) (err error) {
	cpn, ok := target.(*collectnpc.CollectPointNPC)
	if !ok {
		return
	}

	s := cpn.GetScene()
	plMap := s.GetAllPlayers()
	for _, pl := range plMap {
		scMsg := pbutil.BuildSCXiantaoPeachPointChange(cpn)
		pl.SendMsg(scMsg)
	}
	return
}

func init() {
	gameevent.AddEventListener(collecteventtypes.EventTypeCollectPointChange, event.EventListenerFunc(collectPointChange))
}
