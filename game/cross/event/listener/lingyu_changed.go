package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//领域改变
func lingyuChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if !pl.IsCross() {
		return
	}

	playerLingYuChanged := pbutil.BuildPlayerLingYuChanged(pl.GetLingYuId())
	pl.SendCrossMsg(playerLingYuChanged)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowLingYuChanged, event.EventListenerFunc(lingyuChanged))
}
