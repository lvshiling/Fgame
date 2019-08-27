package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	secretcardeventtypes "fgame/fgame/game/secretcard/event/types"
)

//天机牌一键完成
func secretcardFinishAll(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	err = questlogic.FillQuestData(pl, questtypes.QuestSubTypeFinishSecretCard, 0)
	return
}

func init() {
	gameevent.AddEventListener(secretcardeventtypes.EventTypeSecretCardFinishAll, event.EventListenerFunc(secretcardFinishAll))
}
