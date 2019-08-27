package common

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	propertylogic "fgame/fgame/game/property/logic"
	"fgame/fgame/game/scene/scene"
)

//玩家属性变化
func battlePlayerPropertyChanged(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(scene.Player)
	propertylogic.SnapChangedBattleProperty(pl)
	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerPropertyChanged, event.EventListenerFunc(battlePlayerPropertyChanged))
}
