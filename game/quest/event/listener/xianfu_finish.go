package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	xianfueventtypes "fgame/fgame/game/xianfu/event/types"
)

//仙府完成
func xianFuFinsh(target event.EventTarget, data event.EventData) (err error) {

	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData, ok := data.(*xianfueventtypes.XianFuFinishEventData)
	if !ok {
		return
	}
	num := eventData.GetNum()
	if num <= 0 {
		return
	}

	err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeXianFu, 0, num)
	return
}

func init() {
	gameevent.AddEventListener(xianfueventtypes.EventTypeXianFuFinish, event.EventListenerFunc(xianFuFinsh))
}
