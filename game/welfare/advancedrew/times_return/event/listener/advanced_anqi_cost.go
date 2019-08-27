package listener

import (
	"fgame/fgame/core/event"
	anqieventtypes "fgame/fgame/game/anqi/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	advancedrewtimesreturnlogic "fgame/fgame/game/welfare/advancedrew/times_return/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//暗器进阶消耗
func playerAnqiAdvancedCost(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advancedType := welfaretypes.AdvancedTypeAnqi

	//次数返还
	times := int32(1)
	advancedrewtimesreturnlogic.AdvancedTimesReturn(pl, times, advancedType)

	return
}

func init() {
	gameevent.AddEventListener(anqieventtypes.EventTypeAnqiAdvancedCost, event.EventListenerFunc(playerAnqiAdvancedCost))
}
