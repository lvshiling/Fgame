package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	weekeventtypes "fgame/fgame/game/week/event/types"
)

//购买周卡
func playerWeekBuy(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeWeekBuy, 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(weekeventtypes.EventTypeWeekBuy, event.EventListenerFunc(playerWeekBuy))
}
