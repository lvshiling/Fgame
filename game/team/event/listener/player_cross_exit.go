package listener

import (
	"fgame/fgame/core/event"
	crosseventtypes "fgame/fgame/game/cross/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/team/team"
)

//玩家跨服断开
func playerCrossExit(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	teamId := pl.GetTeamId()
	if teamId == 0 {
		return
	}
	team.GetTeamService().ArenaStopMatch(pl)
	return
}

func init() {
	gameevent.AddEventListener(crosseventtypes.EventTypePlayerCrossExit, event.EventListenerFunc(playerCrossExit))
}
