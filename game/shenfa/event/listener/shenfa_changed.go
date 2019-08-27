package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	shenfaeventtypes "fgame/fgame/game/shenfa/event/types"
	playershenfa "fgame/fgame/game/shenfa/player"
)

//身法改变
func shenfaChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	m := pl.GetPlayerDataManager(playertypes.PlayerShenfaDataManagerType).(*playershenfa.PlayerShenfaDataManager)
	shenFaId := m.GetShenFaId()
	pl.SetShenFaId(shenFaId)
	return
}

func init() {
	gameevent.AddEventListener(shenfaeventtypes.EventTypeShenfaChanged, event.EventListenerFunc(shenfaChanged))
}
