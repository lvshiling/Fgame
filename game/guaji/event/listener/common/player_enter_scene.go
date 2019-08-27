package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/guaji/pbutil"
	"fgame/fgame/game/scene/scene"
)

//玩家进入场景
func playerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	s := pl.GetScene()

	if pl.IsGuaJiPlayer() {
		guaJiPos := pbutil.BuildSCGuaJiPos(s.MapId(), pl.GetPosition())
		pl.SendMsg(guaJiPos)
	}

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(playerEnterScene))
}
