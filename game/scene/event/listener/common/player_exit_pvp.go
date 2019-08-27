package common

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"

	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

func playerExitPvp(target event.EventTarget, data event.EventData) (err error) {
	p := target.(scene.Player)

	playerExitPVP := pbutil.BuildPlayerExitPVP(p)
	p.SendMsg(playerExitPVP)
	return nil
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypePlayerExitPVP, event.EventListenerFunc(playerExitPvp))
}
