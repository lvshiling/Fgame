package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
)

//竞技场场景结束
func playerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	// pl, ok := target.(scene.Player)
	// if !ok {
	// 	return
	// }
	// s := pl.GetScene()
	// if s == nil {
	// 	return
	// }
	// if s.MapTemplate().GetMapType() != scenetypes.SceneTypeArenapvp {
	// 	return
	// }
	// scArenaPlayerDataOnlineChanged := pbutil.BuildSCArenapvpPlayerOnlineChanged(pl)
	// s.BroadcastMsg(scArenaPlayerDataOnlineChanged)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(playerEnterScene))
}
