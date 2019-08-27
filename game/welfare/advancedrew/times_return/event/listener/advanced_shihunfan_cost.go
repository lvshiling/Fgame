package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	shihunfaneventtypes "fgame/fgame/game/shihunfan/event/types"
	advancedrewtimesreturnlogic "fgame/fgame/game/welfare/advancedrew/times_return/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//噬魂幡进阶消耗
func playerShiHunFanAdvancedCost(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advancedType := welfaretypes.AdvancedTypeShiHunFan

	//次数返还
	times := int32(1)
	advancedrewtimesreturnlogic.AdvancedTimesReturn(pl, times, advancedType)
	return
}

func init() {
	gameevent.AddEventListener(shihunfaneventtypes.EventTypeShiHunFanAdvancedCost, event.EventListenerFunc(playerShiHunFanAdvancedCost))
}
