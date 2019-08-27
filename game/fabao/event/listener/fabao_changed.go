package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	fabaoeventtypes "fgame/fgame/game/fabao/event/types"
	playerfabao "fgame/fgame/game/fabao/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//法宝改变
func faBaoChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	m := pl.GetPlayerDataManager(playertypes.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	faBaoId := m.GetFaBaoId()
	pl.SetFaBaoId(faBaoId)
	return
}

func init() {
	gameevent.AddEventListener(fabaoeventtypes.EventTypeFaBaoChanged, event.EventListenerFunc(faBaoChanged))
}
