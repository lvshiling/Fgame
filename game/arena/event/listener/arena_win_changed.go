package listener

import (
	"fgame/fgame/core/event"
	arenaeventtypes "fgame/fgame/game/arena/event/types"
	playerarena "fgame/fgame/game/arena/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//竞技场获胜次数变化
func arenaWinChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	arenaManager := pl.GetPlayerDataManager(playertypes.PlayerArenaDataManagerType).(*playerarena.PlayerArenaDataManager)
	arenaObj := arenaManager.GetPlayerArenaObject()
	winCount := arenaObj.GetWinCount()
	pl.SetArenaWinTime(winCount)
	return
}

func init() {
	gameevent.AddEventListener(arenaeventtypes.EventTypeArenaWinChanged, event.EventListenerFunc(arenaWinChanged))
}
