package listener

import (
	"fgame/fgame/core/event"
	additionsyseventtypes "fgame/fgame/game/additionsys/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	smeltlogic "fgame/fgame/game/welfare/drew/smelt/logic"
)

func init() {
	gameevent.AddEventListener(additionsyseventtypes.EventTypeAdditionSysUseItem, event.EventListenerFunc(playerAdditionSysUseItem))
}

func playerAdditionSysUseItem(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	useData, ok := data.(*additionsyseventtypes.PlayerAdditionSysUseItemEventData)
	if !ok {
		return
	}

	return smeltlogic.ListenEventAddItemNum(pl, useData.GetItemMap())
}
