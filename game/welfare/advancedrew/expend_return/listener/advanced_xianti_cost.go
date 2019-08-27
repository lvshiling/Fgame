package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	advancedrewexpendreturnlogic "fgame/fgame/game/welfare/advancedrew/expend_return/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
	xiantieventtypes "fgame/fgame/game/xianti/event/types"
	"fgame/fgame/game/xianti/xianti"
)

//仙体进阶消耗
func playerXianTiAdvancedCost(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advancedId, ok := data.(int32)
	if !ok {
		return
	}

	nexTemp := xianti.GetXianTiService().GetXianTiNumber(advancedId + 1)
	itemNum := nexTemp.ItemCount
	advancedType := welfaretypes.AdvancedTypeXianTi

	//消耗返还（新版）
	advancedrewexpendreturnlogic.AdvancedExpendReturn(pl, itemNum, advancedType)

	return
}

func init() {
	gameevent.AddEventListener(xiantieventtypes.EventTypeXianTiAdvancedCost, event.EventListenerFunc(playerXianTiAdvancedCost))
}
