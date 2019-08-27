package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	advancedblessfeedbacklogic "fgame/fgame/game/welfare/advanced/bless_feedback/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
	wingeventtypes "fgame/fgame/game/wing/event/types"
)

//玩家仙羽进阶
func playerFeatherAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId, ok := data.(int32)
	if !ok {
		return
	}

	advancedType := welfaretypes.AdvancedTypeFeather
	advancedblessfeedbacklogic.UpdateAdvancedBlessActivityData(pl, int32(advanceId), advancedType)
	return
}

func init() {
	gameevent.AddEventListener(wingeventtypes.EventTypeFeatherAdvanced, event.EventListenerFunc(playerFeatherAdavanced))
}
