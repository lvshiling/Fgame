package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	goldequipeventtypes "fgame/fgame/game/goldequip/event/types"
	playergoldequip "fgame/fgame/game/goldequip/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//元神金装镶嵌宝石
func goldEquipEmbedGem(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	err = goldEquipGemTotalLevelChanged(pl)
	if err != nil {
		return
	}
	return nil
}

func goldEquipGemTotalLevelChanged(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	totalLevel := manager.CountTotalGemLevel()
	return questlogic.SetQuestData(pl, questtypes.QuestSubTypeGemTotalLevel, 0, totalLevel)
}

func init() {
	gameevent.AddEventListener(goldequipeventtypes.EventTypeGoldEquipEmbedGem, event.EventListenerFunc(goldEquipEmbedGem))
}
