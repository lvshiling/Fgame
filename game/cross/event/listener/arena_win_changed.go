package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//获胜次数变化
func arenaWinTimeChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if !pl.IsCross() {
		return
	}
	winTime := pl.GetArenaWinTime()
	siArenaPlayerWinTimeChanged := pbutil.BuildSIArenaPlayerWinTimeChanged(winTime)
	pl.SendCrossMsg(siArenaPlayerWinTimeChanged)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerArenaWinTimeChanged, event.EventListenerFunc(arenaWinTimeChanged))
}
