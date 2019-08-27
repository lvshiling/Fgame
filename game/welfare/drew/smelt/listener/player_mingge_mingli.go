package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	minggeeventtypes "fgame/fgame/game/mingge/event/types"
	"fgame/fgame/game/player"
	smeltlogic "fgame/fgame/game/welfare/drew/smelt/logic"
)

func init() {
	gameevent.AddEventListener(minggeeventtypes.EventTypeMingGeMingLi, event.EventListenerFunc(playerMinggeMingli))
}

func playerMinggeMingli(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	useData, ok := data.(*minggeeventtypes.PlayerMingGeMingLiEventData)
	if !ok {
		return
	}

	return smeltlogic.ListenEventAddItemNum(pl, useData.GetUseItemMap())
}
