package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
)

//队伍名称改变
func playerTeamChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	if pl.GetScene() == nil {
		return
	}
	teamName := pl.GetTeamName()
	teamChanged := pbutil.BuildScenePlayerTeamChanged(pl, teamName)
	scenelogic.BroadcastNeighborIncludeSelf(pl, teamChanged)
	return nil
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerTeamChanged, event.EventListenerFunc(playerTeamChanged))
}
