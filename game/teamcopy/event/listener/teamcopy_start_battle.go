package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	teameventtypes "fgame/fgame/game/team/event/types"
	"fgame/fgame/game/team/team"
	teamcopylogic "fgame/fgame/game/teamcopy/logic"
)

//组队副本开始战斗
func teamCopyStartBattle(target event.EventTarget, data event.EventData) (err error) {
	captainPl := target.(player.Player)
	teamObject := data.(*team.TeamObject)

	//推送用户
	teamcopylogic.OnTeamCopyStartBattle(captainPl, teamObject)
	return
}

func init() {
	gameevent.AddEventListener(teameventtypes.EventTypeTeamCopyStartBattle, event.EventListenerFunc(teamCopyStartBattle))
}
