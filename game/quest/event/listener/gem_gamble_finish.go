package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	gemeventtypes "fgame/fgame/game/gem/event/types"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//赌石完成
func gemGambleFinish(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	num := data.(int32)
	if num <= 0 {
		return
	}
	err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeGamble, 0, num)
	return
}

func init() {
	gameevent.AddEventListener(gemeventtypes.EventTypeGemGambleFinish, event.EventListenerFunc(gemGambleFinish))
}
