package listener

import (
	"fgame/fgame/core/event"
	dragoneventtypes "fgame/fgame/game/dragon/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家城战获胜
func playerDragonAdvaced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advancedNum := data.(int32)
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeShenLongFuHua, 0, advancedNum)
	return
}

func init() {
	gameevent.AddEventListener(dragoneventtypes.EventTypeDragonAdvanced, event.EventListenerFunc(playerDragonAdvaced))
}
