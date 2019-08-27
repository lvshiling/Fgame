package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	teamlogic "fgame/fgame/game/team/logic"
	"fgame/fgame/game/team/pbutil"
	"fgame/fgame/game/team/team"
)

//玩家转生变化
func playerZhuanShengChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	playerId := pl.GetId()
	zhuanSheng := data.(int32)

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

	member.SetZhuanSheng(zhuanSheng)
	scTeamBroadcast := pbutil.BuildSCTeamZhuanShengChange(playerId, zhuanSheng)
	teamlogic.BroadcastMsg(teamData, scTeamBroadcast)
	return
}

func init() {
	gameevent.AddEventListener(propertyeventtypes.EventTypePlayerZhuanShengChanged, event.EventListenerFunc(playerZhuanShengChanged))
}
