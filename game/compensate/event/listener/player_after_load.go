package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/compensate/compensate"
	compensatelogic "fgame/fgame/game/compensate/logic"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
)

// 检查补偿
func playerAfterLoad(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	compensateList := compensate.GetCompensateSrevice().GetCompensateList()
	for _, compensateObj := range compensateList {
		compensatelogic.SendServerCompensate(pl, compensateObj)
	}

	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerAfterLoad))
}
