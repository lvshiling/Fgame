package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"

	lingyueventtypes "fgame/fgame/game/lingyu/event/types"
)

//玩家领域进阶
func lingyuAdvanced(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}
	advanceId := data.(int32)
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeSystemX, int32(questtypes.SystemReachXTypeLingYu), advanceId)
	return
}

func init() {
	gameevent.AddEventListener(lingyueventtypes.EventTypeLingyuAdvanced, event.EventListenerFunc(lingyuAdvanced))
}
