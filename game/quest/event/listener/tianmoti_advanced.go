package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	tianmotiventtypes "fgame/fgame/game/tianmo/event/types"
)

//玩家天魔体进阶
func playerTianMoTiAdavancedRew(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId, ok := data.(int32)
	if !ok {
		return
	}
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeSystemX, int32(questtypes.SystemReachXTypeTianMoTi), advanceId)
	return
}

func init() {
	gameevent.AddEventListener(tianmotiventtypes.EventTypeTianMoAdvanced, event.EventListenerFunc(playerTianMoTiAdavancedRew))
}
