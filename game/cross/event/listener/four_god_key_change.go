package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//钥匙改变
func fourGodKeyChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if !pl.IsCross() {
		return
	}

	playerFourGodKeyChanged := pbutil.BuildPlayerFourGodKeyChanged(pl.GetFourGodKey())
	pl.SendCrossMsg(playerFourGodKeyChanged)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowFourGodKeyChanged, event.EventListenerFunc(fourGodKeyChanged))
}
