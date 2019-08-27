package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	playermarry "fgame/fgame/game/marry/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//玩家战力变化
func playerMarryRingChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	ringType := manager.GetRingType()
	pl.SetRingType(ringType)
	return
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypePlayerMarryRingChange, event.EventListenerFunc(playerMarryRingChanged))
}
