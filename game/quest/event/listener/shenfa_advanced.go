package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"

	shenfaeventtypes "fgame/fgame/game/shenfa/event/types"
)

//玩家身法进阶
func shenfaAdvanced(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}

	advanceId := data.(int32)
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeSystemX, int32(questtypes.SystemReachXTypeShenFa), advanceId)
	return
}

func init() {
	gameevent.AddEventListener(shenfaeventtypes.EventTypeShenfaAdvanced, event.EventListenerFunc(shenfaAdvanced))
}
