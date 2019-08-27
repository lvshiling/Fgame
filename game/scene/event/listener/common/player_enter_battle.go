package common

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"

	battleeventtypes "fgame/fgame/game/battle/event/types"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
)

func playerEnterBattle(target event.EventTarget, data event.EventData) (err error) {
	p := target.(scene.Player)
	scenelogic.EnterBattle(p)
	return nil
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypePlayerEnterBattle, event.EventListenerFunc(playerEnterBattle))
}
