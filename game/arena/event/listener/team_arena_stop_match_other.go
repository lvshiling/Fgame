package listener

import (
	"fgame/fgame/core/event"
	arenalogic "fgame/fgame/game/arena/logic"
	crosslogic "fgame/fgame/game/cross/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	teameventtypes "fgame/fgame/game/team/event/types"
	"fgame/fgame/game/team/team"
)

//队伍竞技停止匹配
func teamArenaStopMatchOther(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	teamObj := data.(*team.TeamObject)
	//发送匹配
	captainMem := teamObj.GetCaptain()
	//推送用户
	arenalogic.OnTeamArenaStopMatchOther(pl, teamObj)
	if captainMem.GetPlayerId() == pl.GetId() {
		//退出跨服
		crosslogic.PlayerExitCross(pl)
	} else {
		captainPl := player.GetOnlinePlayerManager().GetPlayerById(captainMem.GetPlayerId())
		if captainPl == nil {
			return
		}
		crosslogic.AsyncPlayerExitCross(captainPl)
	}
	return
}

func init() {
	gameevent.AddEventListener(teameventtypes.EventTypeTeamArenaStopMatchOther, event.EventListenerFunc(teamArenaStopMatchOther))
}
