package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	welfareeventtypes "fgame/fgame/game/welfare/event/types"
)

//玩家完成首冲
func playerFinishFirstCharge(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	err = questlogic.FillQuestData(pl, questtypes.QuestSubTypeFinishFirstCharge, 0)
	return
}

func init() {
	gameevent.AddEventListener(welfareeventtypes.EventTypeFinishFirstCharge, event.EventListenerFunc(playerFinishFirstCharge))
}
