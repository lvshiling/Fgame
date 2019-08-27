package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playerxuechi "fgame/fgame/game/xuechi/player"
)

//从血池补血
func xueChiSync(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerXueChiDataManagerType).(*playerxuechi.PlayerXueChiDataManager)
	manager.Save()
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerXueChiBloodSync, event.EventListenerFunc(xueChiSync))
}
