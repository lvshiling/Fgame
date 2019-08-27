package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	inventoryeventtypes "fgame/fgame/game/inventory/event/types"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//脱下宝石
func equipmentTakeOffGem(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	num := manager.GetEquipGemNum()
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeEmbedGem, 0, num)
	return
}

func init() {
	gameevent.AddEventListener(inventoryeventtypes.EventTypeEquipmentTakeOffGem, event.EventListenerFunc(equipmentTakeOffGem))
}
