package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	goldequipeventtypes "fgame/fgame/game/goldequip/event/types"
	"fgame/fgame/game/player"
)

//卸下元神金装
func goldEquipTakeOff(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	err = embedQualityTwoGoldEquip(pl)
	if err != nil {
		return
	}

	err = embedQualityThreeGoldEquip(pl)
	if err != nil {
		return
	}

	err = embedQualityFourGoldEquip(pl)
	if err != nil {
		return
	}

	err = goldEquipTotalChanged(pl)
	if err != nil {
		return
	}

	return nil
}

func init() {
	gameevent.AddEventListener(goldequipeventtypes.EventTypeGoldEquipTakeOff, event.EventListenerFunc(goldEquipTakeOff))
}
