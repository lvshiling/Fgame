package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//配偶改变
func ringTypeChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if !pl.IsCross() {
		return
	}

	playerRingTypeChanged := pbutil.BuildPlayerRingTypeChanged(pl.GetRingType())
	pl.SendCrossMsg(playerRingTypeChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowRingTypeChanged, event.EventListenerFunc(ringTypeChanged))
}
