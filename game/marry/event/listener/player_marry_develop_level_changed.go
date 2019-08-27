package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/marry/marry"
	"fgame/fgame/game/player"

	marryeventtypes "fgame/fgame/game/marry/event/types"
)

//玩家表白等级变化
func playerMarryDevelopLevelChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	developLevel, ok := data.(int32)
	if !ok {
		return
	}

	pl.SetMarryDevelopLevel(developLevel)
	marry.GetMarryService().SyncMarryDevelopLevel(pl.GetId(), developLevel)
	return
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypePlayerMarryDevelopLevelChanged, event.EventListenerFunc(playerMarryDevelopLevelChanged))
}
