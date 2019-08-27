package listener

import (
	"fgame/fgame/core/event"
	crosslogic "fgame/fgame/game/cross/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	teameventtypes "fgame/fgame/game/team/event/types"
	"fgame/fgame/game/team/team"
	teamcopylogic "fgame/fgame/game/teamcopy/logic"
)

//队伍开始战斗请求失败
func teamCopyStartBattleFailed(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	teamObj := data.(*team.TeamObject)
	//发送匹配
	captainMem := teamObj.GetCaptain()
	//推送用户
	teamcopylogic.OnTeamCopyStartBattleFailed(pl, teamObj)
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
	gameevent.AddEventListener(teameventtypes.EventTypeTeamCopyStartBattleFailed, event.EventListenerFunc(teamCopyStartBattleFailed))
}
