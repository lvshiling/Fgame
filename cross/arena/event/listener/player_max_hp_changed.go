package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/arena/pbutil"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//旗子收集 死亡打断
func playerMaxHPChanged(target event.EventTarget, data event.EventData) (err error) {
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

	scArenaPlayerDataMaxHPChanged := pbutil.BuildSCArenaPlayerDataMaxHPChanged(pl)
	s.BroadcastMsg(scArenaPlayerDataMaxHPChanged)

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerMaxHPChanged, event.EventListenerFunc(playerMaxHPChanged))
}
