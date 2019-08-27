package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

func soulAwakenChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if !pl.IsCross() {
		return
	}

	playerSoulAwakenChanged := pbutil.BuildPlayerSoulAwakenChanged(pl.GetSoulAwakenNum())
	pl.SendCrossMsg(playerSoulAwakenChanged)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerSoulAwakenChanged, event.EventListenerFunc(soulAwakenChanged))
}
