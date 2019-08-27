package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	fabaoventtypes "fgame/fgame/game/fabao/event/types"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家法宝进阶
func playerFaBaoAdavancedRew(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId, ok := data.(int32)
	if !ok {
		return
	}
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeSystemX, int32(questtypes.SystemReachXTypeFaBao), advanceId)
	return
}

func init() {
	gameevent.AddEventListener(fabaoventtypes.EventTypeFaBaoAdvanced, event.EventListenerFunc(playerFaBaoAdavancedRew))
}
