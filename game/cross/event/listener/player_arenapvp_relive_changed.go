package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	"fgame/fgame/game/cross/pbutil"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
)

//竞技场pvp复活次数变化
func arenapvpReliveChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	if !pl.IsCross() {
		return
	}
	reliveTimes := pl.GetArenapvpReliveTimes()
	siMsg := pbutil.BuildSIPlayerArenapvpDataChanged(reliveTimes)
	pl.SendCrossMsg(siMsg)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerArenapvpReliveTimesChanged, event.EventListenerFunc(arenapvpReliveChanged))
}
