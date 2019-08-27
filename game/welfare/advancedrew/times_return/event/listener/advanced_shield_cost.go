package listener

import (
	"fgame/fgame/core/event"
	bodyshieldeventtypes "fgame/fgame/game/bodyshield/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	advancedrewtimesreturnlogic "fgame/fgame/game/welfare/advancedrew/times_return/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//盾刺进阶消耗
func playerShieldAdvancedCost(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advancedType := welfaretypes.AdvancedTypeShield

	//次数返还
	times := int32(1)
	advancedrewtimesreturnlogic.AdvancedTimesReturn(pl, times, advancedType)

	return
}

func init() {
	gameevent.AddEventListener(bodyshieldeventtypes.EventTypeShieldAdvancedCost, event.EventListenerFunc(playerShieldAdvancedCost))
}
