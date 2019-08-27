package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	teamlogic "fgame/fgame/game/team/logic"
	"fgame/fgame/game/team/pbutil"
	"fgame/fgame/game/team/team"
)

//玩家名字变化
func playerNameChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

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

	member.SetName(pl.GetName())
	scTeamBroadcast := pbutil.BuildSCTeamNameChange(pl.GetId(), pl.GetName())
	teamlogic.BroadcastMsg(teamData, scTeamBroadcast)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerNameChanged, event.EventListenerFunc(playerNameChanged))
}
