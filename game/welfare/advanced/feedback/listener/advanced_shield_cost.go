package listener

import (
	"fgame/fgame/core/event"
	"fgame/fgame/game/bodyshield/bodyshield"
	bodyshieldeventtypes "fgame/fgame/game/bodyshield/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	advancedfeedbacklogic "fgame/fgame/game/welfare/advanced/feedback/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//盾刺进阶消耗
func playerShieldAdvancedCost(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advancedId, ok := data.(int32)
	if !ok {
		return
	}
	nexTemp := bodyshield.GetBodyShieldService().GetShield(advancedId + 1)
	itemNum := nexTemp.ItemCount
	advancedType := welfaretypes.AdvancedTypeShield

	//消耗返还（旧版）
	advancedfeedbacklogic.UpdateAdvancedActivityData(pl, itemNum, advancedType)

	return
}

func init() {
	gameevent.AddEventListener(bodyshieldeventtypes.EventTypeShieldAdvancedCost, event.EventListenerFunc(playerShieldAdvancedCost))
}
