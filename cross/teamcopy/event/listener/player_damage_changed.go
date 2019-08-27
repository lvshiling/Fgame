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

//组队副本场景伤害变化
func teamCopySceneDamageChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	sd := data.(teamscene.TeamCopySceneData)
	damage := sd.GetDamage(pl)

	//广播在线成员
	damageChanged := pbutil.BuildSCTeamCopyPlayerDamageChanged(pl.GetId(), damage)
	teamcopylogic.BroadcastTeamCopy(sd, damageChanged)
	return
}

func init() {
	gameevent.AddEventListener(teamcopyeventtypes.EventTypeTeamCopySceneDamageChanged, event.EventListenerFunc(teamCopySceneDamageChanged))
}
