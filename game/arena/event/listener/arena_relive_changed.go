package listener

import (
	"fgame/fgame/core/event"
	arenaeventtypes "fgame/fgame/game/arena/event/types"
	playerarena "fgame/fgame/game/arena/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//竞技场复活次数变化
func arenaReliveChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	arenaManager := pl.GetPlayerDataManager(playertypes.PlayerArenaDataManagerType).(*playerarena.PlayerArenaDataManager)
	arenaObj := arenaManager.GetPlayerArenaObject()
	reliveTime := arenaObj.GetReliveTime()
	pl.SetArenaReliveTime(reliveTime)
	return
}

func init() {
	gameevent.AddEventListener(arenaeventtypes.EventTypeArenaReliveChanged, event.EventListenerFunc(arenaReliveChanged))
}
