package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	teamlogic "fgame/fgame/game/team/logic"
	"fgame/fgame/game/team/pbutil"
	"fgame/fgame/game/team/team"
	teamtypes "fgame/fgame/game/team/types"
)

//玩家下线
func playerLogout(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	playerId := pl.GetId()
	teamId := pl.GetTeamId()

	if teamId == 0 {
		return
	}
	teamData := team.GetTeamService().GetTeam(teamId)
	if teamData == nil {
		return
	}

	member, _ := teamData.GetMember(pl.GetId())
	if member == nil {
		return
	}

	//玩家下线
	levelStatus := team.GetTeamService().PlayerLogout(pl)
	if levelStatus != teamtypes.TeamLeaveStatusTypeDissolve {
		scTeamBroadcast := pbutil.BuildSCTeamOnlineChange(playerId, false)
		teamlogic.BroadcastMsg(teamData, scTeamBroadcast)
	}

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerLogout, event.EventListenerFunc(playerLogout))
}
