package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	shenfaeventtypes "fgame/fgame/game/shenfa/event/types"
	shenfatemplate "fgame/fgame/game/shenfa/template"
	advancedrewexpendreturnlogic "fgame/fgame/game/welfare/advancedrew/expend_return/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//身法进阶消耗
func playerShenfaAdvancedCost(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advancedId, ok := data.(int32)
	if !ok {
		return
	}
	nexTemp := shenfatemplate.GetShenfaTemplateService().GetShenfaByNumber(advancedId + 1)
	itemNum := nexTemp.ItemCount
	advancedType := welfaretypes.AdvancedTypeShenfa

	//消耗返还（新版）
	advancedrewexpendreturnlogic.AdvancedExpendReturn(pl, itemNum, advancedType)
	return
}

func init() {
	gameevent.AddEventListener(shenfaeventtypes.EventTypeShenfaAdvancedCost, event.EventListenerFunc(playerShenfaAdvancedCost))
}
