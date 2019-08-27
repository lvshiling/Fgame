package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/cross/arenapvp/pbutil"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//玩家重生
func playerReborn(target event.EventTarget, data event.EventData) (err error) {
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

	scMsg := pbutil.BuildSCArenapvpPlayerShowDataDeadChanged(pl)
	s.BroadcastMsg(scMsg)
	return
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectReborn, event.EventListenerFunc(playerReborn))
}
