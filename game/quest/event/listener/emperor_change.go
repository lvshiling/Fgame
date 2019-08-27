package listener

import (
	"fgame/fgame/core/event"
	emperoreventtypes "fgame/fgame/game/emperor/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//帝王改变
func emperorChanged(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeEmperor, 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(emperoreventtypes.EmperorEventTypeRobed, event.EventListenerFunc(emperorChanged))
}
