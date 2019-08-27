package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	teameventtypes "fgame/fgame/game/team/event/types"
	"fgame/fgame/game/team/pbutil"
	"fgame/fgame/game/team/team"
)

//申请加入
func teamApplyJoin(target event.EventTarget, data event.EventData) (err error) {
	applyPlayer, ok := target.(player.Player)
	if !ok {
		return
	}
	teamData := data.(*team.TeamObject)
	captain := teamData.GetCaptain()
	captainId := captain.GetPlayerId()
	cpl := player.GetOnlinePlayerManager().GetPlayerById(captainId)
	if cpl == nil {
		return
	}
	applyId := applyPlayer.GetId()
	//推送给队长
	scTeamNearJoinToCaptain := pbutil.BuildSCTeamNearJoinToCaptain(applyId)
	cpl.SendMsg(scTeamNearJoinToCaptain)
	return nil
}

func init() {
	gameevent.AddEventListener(teameventtypes.EventTypeTeamNearApplyJoin, event.EventListenerFunc(teamApplyJoin))
}
