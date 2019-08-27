package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	wingeventtypes "fgame/fgame/game/wing/event/types"
)

//玩家战翼进阶
func playerWingAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId, ok := data.(int)
	if !ok {
		return
	}
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeSystemX, int32(questtypes.SystemReachXTypeWing), int32(advanceId))
	return
}

func init() {
	gameevent.AddEventListener(wingeventtypes.EventTypeWingAdvanced, event.EventListenerFunc(playerWingAdavanced))
}
