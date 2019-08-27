package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	playerdensewat "fgame/fgame/game/densewat/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
)

func denseWatEndTimeSet(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerDenseWatDataManagerType).(*playerdensewat.PlayerDenseWatDataManager)
	manager.SetEndTime()
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerDenseWatEndTimeSet, event.EventListenerFunc(denseWatEndTimeSet))
}
