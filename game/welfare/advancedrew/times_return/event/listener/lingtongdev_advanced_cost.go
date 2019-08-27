package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingtongdeveventtypes "fgame/fgame/game/lingtongdev/event/types"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	"fgame/fgame/game/player"
	advancedrewtimesreturnlogic "fgame/fgame/game/welfare/advancedrew/times_return/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//灵童系统进阶消耗
func lingTongDevAdvancedCost(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	devObj, ok := data.(*playerlingtongdev.PlayerLingTongDevObject)
	if !ok {
		return
	}
	advancedType, ok := welfaretypes.LingTongDevTypeToAdvancedType(devObj.GetClassType())
	if !ok {
		return
	}

	//次数返还
	times := int32(1)
	advancedrewtimesreturnlogic.AdvancedTimesReturn(pl, times, advancedType)
	return
}

func init() {
	gameevent.AddEventListener(lingtongdeveventtypes.EventTypeLingTongDevAdvancedCost, event.EventListenerFunc(lingTongDevAdvancedCost))
}
