package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	inventoryeventtypes "fgame/fgame/game/inventory/event/types"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//装备穿戴
func equipmentPutOn(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	pos, ok := data.(inventorytypes.BodyPositionType)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	equipObj := manager.GetEquipByPos(pos)
	if equipObj == nil {
		return
	}

	itemTemplate := item.GetItemService().GetItem(int(equipObj.ItemId))
	if itemTemplate == nil {
		return
	}
	equipTemplate := itemTemplate.GetEquipmentTemplate()
	if equipTemplate == nil {
		return
	}
	level := equipTemplate.Series
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeEquipmentUpgradeLevel, int32(pos), level)
	return
}

func init() {
	gameevent.AddEventListener(inventoryeventtypes.EventTypeEquipmentPutOn, event.EventListenerFunc(equipmentPutOn))
}
