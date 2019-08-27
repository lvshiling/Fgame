package listener

import (
	"fgame/fgame/core/event"
	chargeeventtypes "fgame/fgame/game/charge/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家完成每日首冲
func playerFinishFirstCycleCharge(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeDayFirstCharge, 0, 1)
	return
}

func init() {
	gameevent.AddEventListener(chargeeventtypes.ChargeEventTypeFirstCycleCharge, event.EventListenerFunc(playerFinishFirstCycleCharge))
}
