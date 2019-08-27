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

//组队副本场景结束
func teamCopySceneEnter(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	sd := data.(teamscene.TeamCopySceneData)

	scTeamCopySceneInfo := pbutil.BuildSCTeamCopySceneInfo(sd)
	pl.SendMsg(scTeamCopySceneInfo)

	scTeamCopyPlayerEnterScene := pbutil.BuildSCTeamCopyPlayerEnterScene(pl)
	teamcopylogic.BroadcastTeamCopyExclude(sd, pl, scTeamCopyPlayerEnterScene)
	return
}

func init() {
	gameevent.AddEventListener(teamcopyeventtypes.EvnetTypeTeamCopySceneEnter, event.EventListenerFunc(teamCopySceneEnter))
}
