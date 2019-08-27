package listener

import (
	"fgame/fgame/core/event"
	arenalogic "fgame/fgame/game/arena/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	teameventtypes "fgame/fgame/game/team/event/types"
	"fgame/fgame/game/team/team"
)

//队伍匹配成功
func teamArenaMatched(target event.EventTarget, data event.EventData) (err error) {

	captainPl := target.(player.Player)
	teamObject := data.(*team.TeamObject)

	//推送匹配成功
	arenalogic.OnTeamArenaMatched(captainPl, teamObject)

	return
}

func init() {
	gameevent.AddEventListener(teameventtypes.EventTypeTeamArenaMatched, event.EventListenerFunc(teamArenaMatched))
}
