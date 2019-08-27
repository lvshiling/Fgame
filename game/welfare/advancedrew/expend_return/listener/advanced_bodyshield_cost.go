package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/bodyshield/bodyshield"
	bodyshieldeventtypes "fgame/fgame/game/bodyshield/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	advancedrewexpendreturnlogic "fgame/fgame/game/welfare/advancedrew/expend_return/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//护体盾进阶消耗
func playerBodyShieldAdvancedCost(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advancedId, ok := data.(int32)
	if !ok {
		return
	}
	nexTemp := bodyshield.GetBodyShieldService().GetBodyShieldNumber(advancedId + 1)
	itemNum := nexTemp.ItemCount
	advancedType := welfaretypes.AdvancedTypeBodyshield

	//消耗返还（新版）
	advancedrewexpendreturnlogic.AdvancedExpendReturn(pl, itemNum, advancedType)

	return
}

func init() {
	gameevent.AddEventListener(bodyshieldeventtypes.EventTypeBodyShieldAdvancedCost, event.EventListenerFunc(playerBodyShieldAdvancedCost))
}
