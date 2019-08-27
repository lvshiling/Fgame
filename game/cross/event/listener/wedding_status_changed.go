package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//结婚状态改变
func weddingStatusChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if !pl.IsCross() {
		return
	}

	playerWeddingStatusChanged := pbutil.BuildPlayerWeddingStatusChanged(pl.GetWeddingStatus())
	pl.SendCrossMsg(playerWeddingStatusChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowWeddingStatusChanged, event.EventListenerFunc(weddingStatusChanged))
}
