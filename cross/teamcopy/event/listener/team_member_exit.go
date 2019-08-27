package listener

import (
	"fgame/fgame/core/event"
	teamcopyeventtypes "fgame/fgame/cross/teamcopy/event/types"
	teamcopylogic "fgame/fgame/cross/teamcopy/logic"
	"fgame/fgame/cross/teamcopy/pbutil"
	teamscene "fgame/fgame/cross/teamcopy/scene"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//玩家退出
func teamCopyMemberExit(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	sd := data.(teamscene.TeamCopySceneData)
	statusChanged := pbutil.BuildSCTeamCopyPlayerStatusChanged(pl.GetId(),  int32(teamscene.MemberStatusGoAway))
	teamcopylogic.BroadcastTeamCopy(sd, statusChanged)
	return
}

func init() {
	gameevent.AddEventListener(teamcopyeventtypes.EventTypeTeamCopyMemberExit, event.EventListenerFunc(teamCopyMemberExit))
}
