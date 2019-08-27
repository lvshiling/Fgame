package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

// 赠送表白道具
func biaoBaiDaoJuGive(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	num, ok := data.(int32)
	if !ok {
		return
	}
	questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeBiaoBaiTimes, 0, num)
	return
}

func init() {
	gameevent.AddEventListener(marryeventtypes.EventTypePlayerMarryBiaoBai, event.EventListenerFunc(biaoBaiDaoJuGive))
}
