package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	goldequipeventtypes "fgame/fgame/game/goldequip/event/types"
	"fgame/fgame/game/player"
	questlogic "fgame/fgame/game/quest/logic"
	questtypes "fgame/fgame/game/quest/types"
)

//元神金装金装分解
func goldEquipResolve(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	num := data.(int32)
	err = questlogic.SetQuestData(pl, questtypes.QuestSubTypeGoldEquipmentResolve, 0, num)
	return nil
}

func init() {
	gameevent.AddEventListener(goldequipeventtypes.EventTypeGoldEquipResolve, event.EventListenerFunc(goldEquipResolve))
}
