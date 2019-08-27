package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	inventoryeventtypes "fgame/fgame/game/inventory/event/types"
	"fgame/fgame/game/player"
	smeltlogic "fgame/fgame/game/welfare/drew/smelt/logic"
)

func init() {
	gameevent.AddEventListener(inventoryeventtypes.EventTypeUseItem, event.EventListenerFunc(playerItemUse))
}

func playerItemUse(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	useData, ok := data.(*inventoryeventtypes.PlayerInventoryItemUseEventData)
	if !ok {
		return
	}

	itemMap := map[int32]int32{
		useData.GetItemId(): useData.GetUseNum(),
	}

	return smeltlogic.ListenEventAddItemNum(pl, itemMap)
}
