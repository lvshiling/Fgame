package listener

import (
	"fgame/fgame/core/event"
	arenaeventtypes "fgame/fgame/game/arena/event/types"
	"fgame/fgame/game/arena/pbutil"
	playerarena "fgame/fgame/game/arena/player"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
)

//3v3积分变化
func playerArenaJiFenChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerArenaDataManagerType).(*playerarena.PlayerArenaDataManager)
	arenaObj := manager.GetPlayerArenaObjectByRefresh()
	scMsg := pbutil.BuildSCPlayerArenaInfo(arenaObj)
	pl.SendMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(arenaeventtypes.EventTypeArenaJiFenChanged, event.EventListenerFunc(playerArenaJiFenChanged))
}
