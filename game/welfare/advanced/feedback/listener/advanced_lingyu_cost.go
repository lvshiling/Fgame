package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingyueventtypes "fgame/fgame/game/lingyu/event/types"
	lingyutemplate "fgame/fgame/game/lingyu/template"
	"fgame/fgame/game/player"
	advancedfeedbacklogic "fgame/fgame/game/welfare/advanced/feedback/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//领域进阶消耗
func playerLingyuAdvancedCost(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advancedId, ok := data.(int32)
	if !ok {
		return
	}
	nexTemp := lingyutemplate.GetLingyuTemplateService().GetLingyuByNumber(advancedId + 1)
	itemNum := nexTemp.ItemCount
	advancedType := welfaretypes.AdvancedTypeLingyu

	//消耗返还（旧版）
	advancedfeedbacklogic.UpdateAdvancedActivityData(pl, itemNum, advancedType)

	return
}

func init() {
	gameevent.AddEventListener(lingyueventtypes.EventTypeLingyuAdvancedCost, event.EventListenerFunc(playerLingyuAdvancedCost))
}
