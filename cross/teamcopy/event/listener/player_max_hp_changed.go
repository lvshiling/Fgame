package listener

import (
	"fgame/fgame/core/event"
	teamcopylogic "fgame/fgame/cross/teamcopy/logic"
	"fgame/fgame/cross/teamcopy/pbutil"
	teamscene "fgame/fgame/cross/teamcopy/scene"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

func playerMaxHPChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}
	switch s.MapTemplate().GetMapType() {
	case scenetypes.SceneTypeCrossTeamCopy:
		{
			sd := s.SceneDelegate().(teamscene.TeamCopySceneData)
			scTeamCopyPlayerMaxHPChanged := pbutil.BuildSCTeamCopyPlayerMaxHPChanged(pl)
			teamcopylogic.BroadcastTeamCopy(sd, scTeamCopyPlayerMaxHPChanged)
		}
	}
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerMaxHPChanged, event.EventListenerFunc(playerMaxHPChanged))
}
