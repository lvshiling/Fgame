package listener

import (
	"fgame/fgame/core/event"
	arenalogic "fgame/fgame/game/arena/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	teameventtypes "fgame/fgame/game/team/event/types"
	"fgame/fgame/game/team/team"
)

//队伍竞技匹配
func teamArenaMatch(target event.EventTarget, data event.EventData) (err error) {
	captainPl := target.(player.Player)
	teamObject := data.(*team.TeamObject)

	//发送匹配
	//推送用户
	arenalogic.OnTeamArenaMatch(captainPl, teamObject)

	return
}

func init() {
	gameevent.AddEventListener(teameventtypes.EventTypeTeamArenaMatch, event.EventListenerFunc(teamArenaMatch))
}
