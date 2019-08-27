package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	relivelogic "fgame/fgame/game/relive/logic"
	"fgame/fgame/game/scene/scene"
)

//复活刷新
func playerReliveRefresh(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(scene.Player)
	if !ok {
		return
	}
	relivelogic.SyncReliveInfo(p)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerReliveRefresh, event.EventListenerFunc(playerReliveRefresh))
}
