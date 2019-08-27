package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/arenapvp/pbutil"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//退出竞技场场景
func playerExit(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	active, ok := data.(bool)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeArenapvp {
		return
	}

	if !active {
		scArenaPlayerDataOfflineChanged := pbutil.BuildSCArenapvpPlayerOfflineChanged(pl)
		s.BroadcastMsg(scArenaPlayerDataOfflineChanged)
	}
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerExitScene, event.EventListenerFunc(playerExit))
}
