package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	fourgodeventtypes "fgame/fgame/game/fourgod/event/types"
	playerfourgod "fgame/fgame/game/fourgod/player"
	inventorylogic "fgame/fgame/game/inventory/logic"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	scenetypes "fgame/fgame/game/scene/types"
)

//玩家四神遗迹获得物品
func fourGodPlayerGetItem(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	s := pl.GetScene()
	if s.MapTemplate().GetMapType() != scenetypes.SceneTypeFourGodWar {
		return
	}

	eventData := data.(*fourgodeventtypes.FourGodItemGetEventData)
	itemId := eventData.GetItemId()
	num := eventData.GetNum()
	manager := pl.GetPlayerDataManager(types.PlayerFourGodDataManagerType).(*playerfourgod.PlayerFourGodDataManager)
	manager.AddItem(itemId, num)
	inventorylogic.SnapInventoryChanged(pl)
	return
}

func init() {
	gameevent.AddEventListener(fourgodeventtypes.EventTypeFourGodGetItem, event.EventListenerFunc(fourGodPlayerGetItem))
}
