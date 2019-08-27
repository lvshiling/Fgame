package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	wardrobeeventtypes "fgame/fgame/game/wardrobe/event/types"
	playerwardrobe "fgame/fgame/game/wardrobe/player"
)

//玩家衣橱套件激活
func playerWardrobeActivate(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	wardrobeObj, ok := data.(*playerwardrobe.PlayerWardrobeObject)
	if !ok {
		return
	}
	wardrobeType := wardrobeObj.GetType()

	err = wardrobeActivateNum(pl, wardrobeType)
	if err != nil {
		return
	}

	err = wardrobeActivateOne(pl)
	if err != nil {
		return
	}

	err = wardrobeActivateTwo(pl)
	if err != nil {
		return
	}

	err = wardrobeActivateThree(pl)
	if err != nil {
		return
	}

	err = wardrobeActivateFour(pl)
	if err != nil {
		return
	}
	return
}

func wardrobeActivateNum(pl player.Player, wardrobeType int32) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	activateNum := manager.GetWardrobeActivateNumByType(wardrobeType)
	if activateNum == 0 {
		return
	}
	err = questlogic.SetQuestDataSurpass(pl, questtypes.QuestSubTypeActivateYiChuSuitNum, int32(wardrobeType), activateNum)
	return
}

func wardrobeActivateOne(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	activateNum := manager.GetWardrobeActivateTypeNumByNum(0)
	if activateNum == 0 {
		return
	}
	err = questlogic.SetQuestDataSurpass(pl, questtypes.QuestSubTypeActivateYiChuSuitOne, 0, activateNum)
	return
}

func wardrobeActivateTwo(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	activateNum := manager.GetWardrobeActivateTypeNumByNum(1)
	if activateNum == 0 {
		return
	}
	err = questlogic.SetQuestDataSurpass(pl, questtypes.QuestSubTypeActivateYiChuSuitTwo, 0, activateNum)
	return
}

func wardrobeActivateThree(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	activateNum := manager.GetWardrobeActivateTypeNumByNum(2)
	if activateNum == 0 {
		return
	}
	err = questlogic.SetQuestDataSurpass(pl, questtypes.QuestSubTypeActivateYiChuSuitThree, 0, activateNum)
	return
}

func wardrobeActivateFour(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerWardrobeDataManagerType).(*playerwardrobe.PlayerWardrobeDataManager)
	activateNum := manager.GetWardrobeActivateTypeNumByNum(3)
	if activateNum == 0 {
		return
	}
	err = questlogic.SetQuestDataSurpass(pl, questtypes.QuestSubTypeActivateYiChuSuitFour, 0, activateNum)
	return
}

func init() {
	gameevent.AddEventListener(wardrobeeventtypes.EventTypeWardrobeActive, event.EventListenerFunc(playerWardrobeActivate))
}
