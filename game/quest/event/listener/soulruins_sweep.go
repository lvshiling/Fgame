package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
	soulruinseventtypes "fgame/fgame/game/soulruins/event/types"
)

//帝陵遗迹扫荡
func soulRuinsSweep(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	num, ok := data.(int32)
	if !ok {
		return
	}
	if num <= 0 {
		return
	}
	err = questlogic.IncreaseQuestData(pl, questtypes.QuestSubTypeEnterSoulRuins, 0, num)
	return
}

func init() {
	gameevent.AddEventListener(soulruinseventtypes.EventTypeSoulruinsSweep, event.EventListenerFunc(soulRuinsSweep))
}
