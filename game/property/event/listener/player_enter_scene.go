package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
)

//玩家玩家进入场景
func playerEnterScene(target event.EventTarget, data event.EventData) (err error) {
	p, ok := target.(player.Player)
	if !ok {
		return
	}
	p.UpdateBattleProperty(playerpropertytypes.PropertyEffectorTypeMaskAll)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerEnterScene, event.EventListenerFunc(playerEnterScene))
}
