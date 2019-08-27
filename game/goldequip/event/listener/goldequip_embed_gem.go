package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	goldequipeventtypes "fgame/fgame/game/goldequip/event/types"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	"fgame/fgame/game/player"
)

//元神金装镶嵌宝石
func goldEquipEmbedGem(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}

	err = goldequiplogic.GoldEquipPropertyChanged(pl)
	return
}

func init() {
	gameevent.AddEventListener(goldequipeventtypes.EventTypeGoldEquipEmbedGem, event.EventListenerFunc(goldEquipEmbedGem))
}
