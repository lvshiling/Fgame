package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	smeltlogic "fgame/fgame/game/welfare/drew/smelt/logic"
	wushuangweaponeventtypes "fgame/fgame/game/wushuangweapon/event/types"
)

func init() {
	gameevent.AddEventListener(wushuangweaponeventtypes.EventTypeWushuangDevouring, event.EventListenerFunc(playerWushuangDevouring))
}

func playerWushuangDevouring(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	useData, ok := data.(*wushuangweaponeventtypes.PlayerWushuangDevouringEventData)
	if !ok {
		return
	}

	return smeltlogic.ListenEventAddItemNum(pl, useData.GetUseItemMap())
}
