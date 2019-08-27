package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	jieyieventtypes "fgame/fgame/game/jieyi/event/types"
	playerjieyi "fgame/fgame/game/jieyi/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//结义改变
func jieYiChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	m := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	jieYiId := m.GetJieYiId()
	jieYiName := m.GetJieYiName()
	jieYiRank := m.GetJieYiRank()

	pl.SyncJieYi(jieYiId, jieYiName, jieYiRank)
	return
}

func init() {
	gameevent.AddEventListener(jieyieventtypes.JieYiEventTypeJieYiChange, event.EventListenerFunc(jieYiChanged))
}
