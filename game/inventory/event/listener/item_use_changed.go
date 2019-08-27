package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	inventoryeventtypes "fgame/fgame/game/inventory/event/types"
	"fgame/fgame/game/inventory/pbutil"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
)

//物品使用信息变化
func itemUseChanged(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	itemUseMap := data.(map[int32]*playerinventory.PlayerItemUseObject)

	//推送所有物品
	scInventoryItemUseChangedNotice := pbutil.BuildSCInventoryItemUseChangedNotice(itemUseMap)
	p.SendMsg(scInventoryItemUseChangedNotice)
	return
}

func init() {
	gameevent.AddEventListener(inventoryeventtypes.EventTypeItemUseChanged, event.EventListenerFunc(itemUseChanged))
}
