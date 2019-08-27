package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/xuechi/pbutil"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//从血池补血
func xueChiRecover(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}

	isXueChiSync := pbutil.BuildISXueChiSync(pl)
	pl.SendMsg(isXueChiSync)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerXueChiRecover, event.EventListenerFunc(xueChiRecover))
}
