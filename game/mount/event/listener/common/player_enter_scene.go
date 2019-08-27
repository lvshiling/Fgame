package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/scene/scene"
)

//玩家玩家进入场景
func playerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	p := target.(scene.Player)
	s := p.GetScene()
	limitRide := s.MapTemplate().LimitRideHorse
	if limitRide == 0 {
		return
	}
	if p.IsMountHidden() {
		return
	}
	p.MountHidden(true)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(playerEnterScene))
}
