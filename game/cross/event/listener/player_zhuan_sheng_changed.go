package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

func playerZhuanShengChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if !pl.IsCross() {
		return
	}

	playerZhuanShengChanged := pbutil.BuildPlayerZhuanShengChanged(pl.GetZhuanSheng())
	pl.SendCrossMsg(playerZhuanShengChanged)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerZhuanShengChanged, event.EventListenerFunc(playerZhuanShengChanged))
}
