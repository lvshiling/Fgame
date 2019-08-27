package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//灵童变化
func playerLingTongShow(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}
	lingTong := pl.GetLingTong()
	if lingTong == nil {
		return
	}
	if pl.IsLingTongHidden() {
		s.RemoveSceneObject(lingTong, true)
	} else {
		s.AddSceneObject(lingTong)
	}

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerLingTongShow, event.EventListenerFunc(playerLingTongShow))
}
