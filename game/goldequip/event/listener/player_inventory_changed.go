package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	inventoryeventtypes "fgame/fgame/game/inventory/event/types"
	"fgame/fgame/game/player"
)

//玩家物品变更
func playerInventoryChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	//自动分解
	goldequiplogic.AutoFenJieGoldEquip(pl)
	return
}

func init() {
	gameevent.AddEventListener(inventoryeventtypes.EventTypeInventoryChanged, event.EventListenerFunc(playerInventoryChanged))
}
