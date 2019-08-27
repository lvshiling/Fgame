package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	teameventtypes "fgame/fgame/game/team/event/types"
	"fgame/fgame/game/team/team"
	teamcopylogic "fgame/fgame/game/teamcopy/logic"
)

//队伍开始战斗请求成功
func teamCopyStartBattleSucess(target event.EventTarget, data event.EventData) (err error) {
	captainPl := target.(player.Player)
	teamObject := data.(*team.TeamObject)

	teamcopylogic.OnTeamCopyStartBattleSucess(captainPl, teamObject)
	return
}

func init() {
	gameevent.AddEventListener(teameventtypes.EventTypeTeamCopyStartBattleSucess, event.EventListenerFunc(teamCopyStartBattleSucess))
}
