package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	massacreventtypes "fgame/fgame/game/massacre/event/types"
	massacretemplate "fgame/fgame/game/massacre/template"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//玩家戮仙刃进阶
func playerMassacreAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	eventData := data.(*massacreventtypes.PlayerMassacreAdvancedEventData)
	advanceId := eventData.GetNewAdvanceId()
	massacreTemplate := massacretemplate.GetMassacreTemplateService().GetMassacre(int(advanceId))
	if massacreTemplate == nil {
		return
	}
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeSystemX, int32(questtypes.SystemReachXTypeLuXianRen), massacreTemplate.Type)
	return
}

func init() {
	gameevent.AddEventListener(massacreventtypes.EventTypeMassacreAdvanced, event.EventListenerFunc(playerMassacreAdavanced))
}
