package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//法宝变化
func faBaoChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if !pl.IsCross() {
		return
	}

	playerFaBaoChanged := pbutil.BuildPlayerFaBaoChanged(pl.GetFaBaoId())
	pl.SendCrossMsg(playerFaBaoChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowFaBaoChanged, event.EventListenerFunc(faBaoChanged))
}
