package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	shihunfaneventtypes "fgame/fgame/game/shihunfan/event/types"
)

//玩家噬魂幡进阶
func shiHunFanAdvanced(target event.EventTarget, data event.EventData) (err error) {
	pl := target.(player.Player)
	if pl == nil {
		return
	}

	advanceId := data.(int32)
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeSystemX, int32(questtypes.SystemReachXTypeShiHunFan), advanceId)
	return
}

func init() {
	gameevent.AddEventListener(shihunfaneventtypes.EventTypeShiHunFanAdvanced, event.EventListenerFunc(shiHunFanAdvanced))
}
