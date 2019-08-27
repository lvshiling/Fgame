package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	tulongequipeventtypes "fgame/fgame/game/tulongequip/event/types"
	smeltlogic "fgame/fgame/game/welfare/drew/smelt/logic"
)

func init() {
	gameevent.AddEventListener(tulongequipeventtypes.EventTypeTuLongEquipUseItem, event.EventListenerFunc(playerTulongEquipUseItem))
}

func playerTulongEquipUseItem(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	useData, ok := data.(*tulongequipeventtypes.PlayerTuLongEquipUseItemEventData)
	if !ok {
		return
	}

	return smeltlogic.ListenEventAddItemNum(pl, useData.GetUseItemMap())
}
