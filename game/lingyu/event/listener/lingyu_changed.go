package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingyueventtypes "fgame/fgame/game/lingyu/event/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"

	playerlingyu "fgame/fgame/game/lingyu/player"
)

//领域改变
func lingyuChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	m := pl.GetPlayerDataManager(playertypes.PlayerLingyuDataManagerType).(*playerlingyu.PlayerLingyuDataManager)
	lingYuId := m.GetLingYuId()
	pl.SetLingYuId(lingYuId)
	return
}

func init() {
	gameevent.AddEventListener(lingyueventtypes.EventTypeLingyuChanged, event.EventListenerFunc(lingyuChanged))
}
