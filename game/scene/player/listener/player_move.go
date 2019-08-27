package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
)

//玩家玩家进入场景
func playerMove(target event.EventTarget, data event.EventData) (err error) {

	return
}

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerMove, event.EventListenerFunc(playerMove))
}
