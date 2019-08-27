package common

import (
	"fgame/fgame/core/event"
	coretypes "fgame/fgame/core/types"
	gameevent "fgame/fgame/game/event"
	sceneeventtypes "fgame/fgame/game/scene/event/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

//玩家复活
func battleObjectReborn(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)
	pos := data.(coretypes.Position)
	scenelogic.OnReborn(bo, pos)

	return nil
}

func init() {
	gameevent.AddEventListener(sceneeventtypes.EventTypeBattleObjectReborn, event.EventListenerFunc(battleObjectReborn))
}
