package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	souleventtypes "fgame/fgame/game/soul/event/types"
	playersoul "fgame/fgame/game/soul/player"
	soultypes "fgame/fgame/game/soul/types"
)

//帝魂激活
func soulActive(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	soulTag, ok := data.(soultypes.SoulType)
	if !ok {
		return
	}
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeSoulActive, int32(soulTag), 1)
	if err != nil {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)
	soulNum := manager.GetSoulNum()
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeActivateSoulNum, 0, soulNum)
	return
}

func init() {
	gameevent.AddEventListener(souleventtypes.EventTypeSoulActive, event.EventListenerFunc(soulActive))
}
