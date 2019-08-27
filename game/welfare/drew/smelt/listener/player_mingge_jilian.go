package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	minggeeventtypes "fgame/fgame/game/mingge/event/types"
	"fgame/fgame/game/player"
	smeltlogic "fgame/fgame/game/welfare/drew/smelt/logic"
)

func init() {
	gameevent.AddEventListener(minggeeventtypes.EventTypeMingGeJiLian, event.EventListenerFunc(playerMinggeJilian))
}

func playerMinggeJilian(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	useData, ok := data.(*minggeeventtypes.PlayerMingGeJiLianEventData)
	if !ok {
		return
	}

	return smeltlogic.ListenEventAddItemNum(pl, useData.GetUseItemMap())
}
