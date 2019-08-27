package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//身法改变
func shenfaChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if !pl.IsCross() {
		return
	}

	playerShenFaChanged := pbutil.BuildPlayerShenFaChanged(pl.GetShenFaId())
	pl.SendCrossMsg(playerShenFaChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowShenFaChanged, event.EventListenerFunc(shenfaChanged))
}
