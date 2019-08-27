package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	teamlogic "fgame/fgame/game/team/logic"
	"fgame/fgame/game/team/pbutil"
	"fgame/fgame/game/team/team"
)

//玩家战力变化
func playerForceChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	teamId := pl.GetTeamId()
	if teamId == 0 {
		return
	}
	force := pl.GetForce()
	teamData := team.GetTeamService().UpdateMemberForce(pl.GetId(), force)
	if teamData == nil {
		return
	}
	playerId := pl.GetId()
	scTeamBroadcast := pbutil.BuildSCTeamForceChange(playerId, force)
	teamlogic.BroadcastMsg(teamData, scTeamBroadcast)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerForceChanged, event.EventListenerFunc(playerForceChanged))
}
