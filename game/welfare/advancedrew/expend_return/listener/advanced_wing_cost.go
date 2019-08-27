package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	advancedrewexpendreturnlogic "fgame/fgame/game/welfare/advancedrew/expend_return/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
	wingeventtypes "fgame/fgame/game/wing/event/types"
	"fgame/fgame/game/wing/wing"
)

//战翼进阶消耗
func playerWingAdvancedCost(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advancedId, ok := data.(int32)
	if !ok {
		return
	}
	nexTemp := wing.GetWingService().GetWingNumber(advancedId + 1)
	itemNum := nexTemp.ItemCount
	advancedType := welfaretypes.AdvancedTypeWing

	//消耗返还（新版）
	advancedrewexpendreturnlogic.AdvancedExpendReturn(pl, itemNum, advancedType)

	return
}

func init() {
	gameevent.AddEventListener(wingeventtypes.EventTypeWingAdvancedCost, event.EventListenerFunc(playerWingAdvancedCost))
}
