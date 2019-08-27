package listener

import (
	"fgame/fgame/core/event"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	gameevent "fgame/fgame/game/event"
)

func init() {
	gameevent.AddEventListener(battleeventtypes.EventTypeBattlePlayerShowMountSync, event.EventListenerFunc(mountHidden))
}
