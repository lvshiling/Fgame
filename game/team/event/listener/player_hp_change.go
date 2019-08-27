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

//玩家血量变化
func playerHpChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	playerId := pl.GetId()
	hp := pl.GetHP()

	teamId := pl.GetTeamId()
	if teamId == 0 {
		return
	}
	teamData := team.GetTeamService().GetTeam(teamId)
	if teamData == nil {
		return
	}

	mem, _ := teamData.GetMember(pl.GetId())
	if mem == nil {
		return
	}

	scTeamBroadcast := pbutil.BuildSCTeamHpChange(playerId, hp)
	teamlogic.BroadcastMsg(teamData, scTeamBroadcast)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerHPChanged, event.EventListenerFunc(playerHpChanged))
}
