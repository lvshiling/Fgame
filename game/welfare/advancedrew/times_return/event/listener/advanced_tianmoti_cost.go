package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	tianmoeventtypes "fgame/fgame/game/tianmo/event/types"
	advancedrewtimesreturnlogic "fgame/fgame/game/welfare/advancedrew/times_return/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//天魔体进阶消耗
func playerTianMoTiAdvancedCost(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advancedType := welfaretypes.AdvancedTypeTianMoTi

	//次数返还
	times := int32(1)
	advancedrewtimesreturnlogic.AdvancedTimesReturn(pl, times, advancedType)
	return
}

func init() {
	gameevent.AddEventListener(tianmoeventtypes.EventTypeTianMoAdvancedCost, event.EventListenerFunc(playerTianMoTiAdvancedCost))
}
