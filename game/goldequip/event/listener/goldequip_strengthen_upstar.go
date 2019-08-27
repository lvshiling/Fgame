package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	goldequipeventtypes "fgame/fgame/game/goldequip/event/types"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	"fgame/fgame/game/player"
)

//玩家元神金装升星强化
func playerGoldEquipStrengUpstar(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	goldequiplogic.GoldEquipPropertyChanged(pl)
	return
}

func init() {
	gameevent.AddEventListener(goldequipeventtypes.EventTypeGoldEquipStrengUpstarSuccess, event.EventListenerFunc(playerGoldEquipStrengUpstar))
}
