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

//装备强化
func equipmentStrengthen(target event.EventTarget, data event.EventData) (err error) {
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
	level := equipSoltObj.Level
	totalLevel := manager.GetEquipTotalLevel()

	err = equipmentStrengthenLevel(pl, int32(pos), level)
	if err != nil {
		return
	}

	err = equipmentStrengthenTotalLevel(pl, totalLevel)
	if err != nil {
		return
	}
	return
}

//指定部位装备强化等级为X级
func equipmentStrengthenLevel(pl player.Player, pos int32, level int32) (err error) {
	return questlogic.SetQuestData(pl, questtypes.QuestSubTypeEquipmentStrengthenLevel, pos, level)
}

//强化总等级为X级
func equipmentStrengthenTotalLevel(pl player.Player, totalLevel int32) (err error) {
	return questlogic.SetQuestData(pl, questtypes.QuestSubTypeEquipmentStrengthenTotalLevel, 0, totalLevel)
}

func init() {
	gameevent.AddEventListener(inventoryeventtypes.EventTypeEquipmentStrengthenLevel, event.EventListenerFunc(equipmentStrengthen))
}
