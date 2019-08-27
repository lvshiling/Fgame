package common

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/battle/battle"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//灵童变化
func playerLingTongChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(scene.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s == nil {
		return
	}

	if pl.IsLingTongHidden() {
		return
	}
	eventData := data.(*battle.LingTongChangeEventData)

	oldLingTong := eventData.GetOldLingTong()
	if oldLingTong != nil {
		s.RemoveSceneObject(oldLingTong, true)
	}
	newLingTong := eventData.GetNewLingTong()
	if newLingTong != nil {
		s.AddSceneObject(newLingTong)
	}
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerLingTongChanged, event.EventListenerFunc(playerLingTongChanged))
}
