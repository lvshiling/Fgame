package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	playermarry "fgame/fgame/game/marry/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//玩家结婚状态变化
func playerMarryWedStatusChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(playertypes.PlayerMarryDataManagerType).(*playermarry.PlayerMarryDataManager)
	weddingStatus := int32(manager.GetWedStatus())
	pl.SetWeddingStatus(weddingStatus)
	return
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypePlayerMarryWedStatusChange, event.EventListenerFunc(playerMarryWedStatusChanged))
}
