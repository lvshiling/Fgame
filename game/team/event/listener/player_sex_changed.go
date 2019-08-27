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

//玩家性别变化
func playerSexChanged(target event.EventTarget, data event.EventData) (err error) {
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

	member.SetSex(pl.GetSex())

	scTeamBroadcast := pbutil.BuildSCTeamSexChange(pl.GetId(), pl.GetSex())
	teamlogic.BroadcastMsg(teamData, scTeamBroadcast)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerSexChanged, event.EventListenerFunc(playerSexChanged))
}
