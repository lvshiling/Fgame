package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	shihunfaneventtypes "fgame/fgame/game/shihunfan/event/types"
	shihunfantemplate "fgame/fgame/game/shihunfan/template"
	advancedrewexpendreturnlogic "fgame/fgame/game/welfare/advancedrew/expend_return/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//噬魂幡进阶消耗
func playerShiHunFanAdvancedCost(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advancedId, ok := data.(int32)
	if !ok {
		return
	}
	nexTemp := shihunfantemplate.GetShiHunFanTemplateService().GetShiHunFanNumber(advancedId + 1)
	itemNum := nexTemp.ItemCount
	advancedType := welfaretypes.AdvancedTypeShiHunFan
	//消耗返还（新版）
	advancedrewexpendreturnlogic.AdvancedExpendReturn(pl, itemNum, advancedType)
	return
}

func init() {
	gameevent.AddEventListener(shihunfaneventtypes.EventTypeShiHunFanAdvancedCost, event.EventListenerFunc(playerShiHunFanAdvancedCost))
}
