package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	fabaoeventtypes "fgame/fgame/game/fabao/event/types"
	fabaotemplate "fgame/fgame/game/fabao/template"
	"fgame/fgame/game/player"
	advancedrewexpendreturnlogic "fgame/fgame/game/welfare/advancedrew/expend_return/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//法宝进阶消耗
func playerFaBaoAdvancedCost(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advancedId, ok := data.(int32)
	if !ok {
		return
	}

	nexTemp := fabaotemplate.GetFaBaoTemplateService().GetFaBaoNumber(advancedId + 1)
	itemNum := nexTemp.ItemCount
	advancedType := welfaretypes.AdvancedTypeFaBao

	//消耗返还（新版）
	advancedrewexpendreturnlogic.AdvancedExpendReturn(pl, itemNum, advancedType)

	return
}

func init() {
	gameevent.AddEventListener(fabaoeventtypes.EventTypeFaBaoAdvancedCost, event.EventListenerFunc(playerFaBaoAdvancedCost))
}
