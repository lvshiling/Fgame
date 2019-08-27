package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	goldequipeventtypes "fgame/fgame/game/goldequip/event/types"
	"fgame/fgame/game/player"
)

//元神金装一键镶嵌
func goldEquipUseGemAll(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	err = goldEquipGemTotalLevelChanged(pl)
	if err != nil {
		return
	}
	return nil
}

func init() {
	gameevent.AddEventListener(goldequipeventtypes.EventTypeGoldEquipUseGemAll, event.EventListenerFunc(goldEquipUseGemAll))
}
