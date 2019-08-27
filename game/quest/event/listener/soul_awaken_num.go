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
)

//帝魂觉醒
func playerSoulAwaken(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)
	awakenNum := manager.GetAwakenNum()
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeAwakenSoulNum, 0, awakenNum)
	return
}

func init() {
	gameevent.AddEventListener(souleventtypes.EventTypeSoulAwaken, event.EventListenerFunc(playerSoulAwaken))
}
