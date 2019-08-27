package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	xiantieventtypes "fgame/fgame/game/xianti/event/types"
)

//玩家仙体进阶
func playerXianTiAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advancedId, ok := data.(int)
	if !ok {
		return
	}
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeSystemX, int32(questtypes.SystemReachXTypeXianTi), int32(advancedId))
	return
}

func init() {
	gameevent.AddEventListener(xiantieventtypes.EventTypeXianTiAdvanced, event.EventListenerFunc(playerXianTiAdavanced))
}
