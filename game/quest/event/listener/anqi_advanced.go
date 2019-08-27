package listener

import (
	"fgame/fgame/core/event"
	anqiventtypes "fgame/fgame/game/anqi/event/types"
	anqitemplate "fgame/fgame/game/anqi/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家暗器进阶
func playerAnqiAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId := data.(int32)
	anqiTemplate := anqitemplate.GetAnqiTemplateService().GetAnqiNumber(advanceId)
	if anqiTemplate == nil {
		return
	}
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeSystemX, int32(questtypes.SystemReachXTypeAnQi), advanceId)
	return
}

func init() {
	gameevent.AddEventListener(anqiventtypes.EventTypeAnqiAdvanced, event.EventListenerFunc(playerAnqiAdavanced))
}
