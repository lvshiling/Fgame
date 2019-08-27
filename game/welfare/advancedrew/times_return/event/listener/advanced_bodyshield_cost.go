package listener

import (
	"fgame/fgame/core/event"
	bodyshieldeventtypes "fgame/fgame/game/bodyshield/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	advancedrewtimesreturnlogic "fgame/fgame/game/welfare/advancedrew/times_return/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//护体盾进阶消耗
func playerBodyShieldAdvancedCost(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advancedType := welfaretypes.AdvancedTypeBodyshield

	//次数返还
	times := int32(1)
	advancedrewtimesreturnlogic.AdvancedTimesReturn(pl, times, advancedType)

	return
}

func init() {
	gameevent.AddEventListener(bodyshieldeventtypes.EventTypeBodyShieldAdvancedCost, event.EventListenerFunc(playerBodyShieldAdvancedCost))
}
