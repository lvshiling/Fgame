package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	xiantieventtypes "fgame/fgame/game/xianti/event/types"
	playerxianti "fgame/fgame/game/xianti/player"
)

//仙体改变
func xianTiChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	m := pl.GetPlayerDataManager(playertypes.PlayerXianTiDataManagerType).(*playerxianti.PlayerXianTiDataManager)
	xianTiId := m.GetXianTiId()
	pl.SetXianTiId(xianTiId)
	return
}

func init() {
	gameevent.AddEventListener(xiantieventtypes.EventTypeXianTiChanged, event.EventListenerFunc(xianTiChanged))
}
