package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//竞技场复活次数变化
func arenaReliveChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if !pl.IsCross() {
		return
	}
	reliveTime := pl.GetArenaReliveTime()
	siArenaPlayerReliveTimeChanged := pbutil.BuildSIArenaPlayerReliveTimeChanged(reliveTime)
	pl.SendCrossMsg(siArenaPlayerReliveTimeChanged)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerArenaReliveTimeChanged, event.EventListenerFunc(arenaReliveChanged))
}
