package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/arenapvp/pbutil"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//最大血量变化
func playerMaxHPChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
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
	scMsg := pbutil.BuildSCArenapvpPlayerShowDataMaxHpChanged(pl)
	s.BroadcastMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerMaxHPChanged, event.EventListenerFunc(playerMaxHPChanged))
}
