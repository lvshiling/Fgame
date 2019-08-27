package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	advancedrewtimesreturnlogic "fgame/fgame/game/welfare/advancedrew/times_return/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
	xiantieventtypes "fgame/fgame/game/xianti/event/types"
)

//仙体进阶消耗
func playerXianTiAdvancedCost(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	
	//次数返还
	advancedType := welfaretypes.AdvancedTypeXianTi
	times := int32(1)
	advancedrewtimesreturnlogic.AdvancedTimesReturn(pl, times, advancedType)

	return
}

func init() {
	gameevent.AddEventListener(xiantieventtypes.EventTypeXianTiAdvancedCost, event.EventListenerFunc(playerXianTiAdvancedCost))
}
