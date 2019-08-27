package listener

import (
	"fgame/fgame/core/event"
	teamcopyeventtypes "fgame/fgame/cross/teamcopy/event/types"
	teamcopylogic "fgame/fgame/cross/teamcopy/logic"
	"fgame/fgame/cross/teamcopy/scene"
	gameevent "fgame/fgame/game/event"
)

//组队副本场景创建成功
func teamCopySceneCreateFinish(target event.EventTarget, data event.EventData) (err error) {
	sd := target.(scene.TeamCopySceneData)
	teamcopylogic.OnTeamCopyRobotEnterScene(sd)
	return
}

func init() {
	gameevent.AddEventListener(teamcopyeventtypes.EventTypeTeamCopySceneCreateFinish, event.EventListenerFunc(teamCopySceneCreateFinish))
}
