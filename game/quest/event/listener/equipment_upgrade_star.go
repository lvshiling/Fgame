package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	inventoryeventtypes "fgame/fgame/game/inventory/event/types"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//装备升星
func equipmentUpgradeStar(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	pos, ok := data.(inventorytypes.BodyPositionType)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	equipSoltObj := manager.GetEquipByPos(pos)
	if equipSoltObj == nil {
		return
	}
	equipStar := equipSoltObj.Star
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeEquipmentUpgradeStar, int32(pos), equipStar)
	return
}

func init() {
	gameevent.AddEventListener(inventoryeventtypes.EventTypeEquipmentUpgradeStar, event.EventListenerFunc(equipmentUpgradeStar))
}
