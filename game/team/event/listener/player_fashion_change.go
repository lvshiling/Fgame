package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	fashioneventtypes "fgame/fgame/game/fashion/event/types"
	"fgame/fgame/game/player"
	teamlogic "fgame/fgame/game/team/logic"
	"fgame/fgame/game/team/pbutil"
	"fgame/fgame/game/team/team"
)

//玩家时装变化变化
func fashionChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	teamId := pl.GetTeamId()
	if teamId == 0 {
		return
	}

	playerId := pl.GetId()
	fashionId := pl.GetFashionId()
	teamData := team.GetTeamService().GetTeam(teamId)
	if teamData == nil {
		return
	}
	member, _ := teamData.GetMember(playerId)
	if member == nil {
		return
	}

	member.SetFashionId(fashionId)
	scTeamBroadcast := pbutil.BuildSCTeamFashionChange(playerId, fashionId)
	teamlogic.BroadcastMsg(teamData, scTeamBroadcast)
	return
}

func init() {
	gameevent.AddEventListener(fashioneventtypes.EventTypeFashionChanged, event.EventListenerFunc(fashionChanged))
}
