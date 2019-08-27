package common

import (
	"fgame/fgame/core/event"
	coretypes "fgame/fgame/core/types"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

//玩家复活
func battleObjectMoveFix(target event.EventTarget, data event.EventData) (err error) {
	bo := target.(scene.BattleObject)

	destPos := data.(coretypes.Position)

	scenelogic.FixPosition(bo, destPos)

	return nil
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattleObjectMoveFix, event.EventListenerFunc(battleObjectMoveFix))
}
