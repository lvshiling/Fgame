package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/arena/pbutil"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//旗子收集 死亡打断
func playerDead(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeArena {
		return
	}

	scArenaPlayerDataDeadChanged := pbutil.BuildSCArenaPlayerDataDeadChanged(pl)
	s.BroadcastMsg(scArenaPlayerDataDeadChanged)

	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectDead, event.EventListenerFunc(playerDead))
}
