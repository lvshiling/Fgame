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

//穿戴元神金装品质大于
func goldEquipPutOn(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	err = embedQualityTwoGoldEquip(pl)
	if err != nil {
		return
	}

	err = embedQualityThreeGoldEquip(pl)
	if err != nil {
		return
	}

	err = embedQualityFourGoldEquip(pl)
	if err != nil {
		return
	}

	err = goldEquipTotalChanged(pl)
	if err != nil {
		return
	}


	return err
}

//穿戴品质达到2及以上的元神金装x件()
func embedQualityTwoGoldEquip(pl player.Player) (err error) {
	return questlogic.SetQuestEmbedData(pl, questtypes.QuestSubTypeEmbedQualityTwoGoldEquipNum)
}

//穿戴品质达到3及以上的元神金装x件()
func embedQualityThreeGoldEquip(pl player.Player) (err error) {
	return questlogic.SetQuestEmbedData(pl, questtypes.QuestSubTypeEmbedQualityThreeGoldEquipNum)
}

//穿戴品质达到4及以上的元神金装x件()
func embedQualityFourGoldEquip(pl player.Player) (err error) {
	return questlogic.SetQuestEmbedData(pl, questtypes.QuestSubTypeEmbedQualityFourGoldEquipNum)
}

func goldEquipTotalChanged(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	totalLevel := manager.CountTotalUpstarLevel()
	return questlogic.SetQuestData(pl, questtypes.QuestSubTypeGoldEquipmentTotalLevel, 0, totalLevel)
}



func init() {
	gameevent.AddEventListener(goldequipeventtypes.EventTypeGoldEquipPutOn, event.EventListenerFunc(goldEquipPutOn))
}
