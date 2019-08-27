package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/arena/arena"
	playereventtypes "fgame/fgame/cross/player/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	teamtypes "fgame/fgame/game/team/types"
)

//跨服登陆
func crossPlayerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	//TODO 修改
	pl := target.(scene.Player)
	arenaTeam := arena.GetArenaService().GetArenaTeamByPlayerId(pl.GetId())
	if arenaTeam == nil {
		return
	}

	arena.GetArenaService().ArenaMemberOnline(pl)
	pl.SetArenaTeam(arenaTeam.GetTeam().GetTeamId(), arenaTeam.GetTeam().GetTeamName(), teamtypes.TeamPurposeTypeNormal)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypeCrossPlayerAfterLoad, event.EventListenerFunc(crossPlayerAfterLoad))
}
