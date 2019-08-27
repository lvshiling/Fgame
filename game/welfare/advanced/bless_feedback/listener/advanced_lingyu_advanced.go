package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	lingyuventtypes "fgame/fgame/game/lingyu/event/types"
	"fgame/fgame/game/player"
	advancedblessfeedbacklogic "fgame/fgame/game/welfare/advanced/bless_feedback/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家领域进阶
func playerLingyuAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId, ok := data.(int32)
	if !ok {
		return
	}

	advancedType := welfaretypes.AdvancedTypeLingyu
	advancedblessfeedbacklogic.UpdateAdvancedBlessActivityData(pl, advanceId, advancedType)
	return
}

func init() {
	gameevent.AddEventListener(lingyuventtypes.EventTypeLingyuAdvanced, event.EventListenerFunc(playerLingyuAdavanced))
}