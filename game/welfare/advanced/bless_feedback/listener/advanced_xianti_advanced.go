package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	advancedblessfeedbacklogic "fgame/fgame/game/welfare/advanced/bless_feedback/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
	xiantiventtypes "fgame/fgame/game/xianti/event/types"
)

//玩家仙体进阶
func playerXianTiAdavancedRew(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId, ok := data.(int)
	if !ok {
		return
	}

	advancedType := welfaretypes.AdvancedTypeXianTi
	advancedblessfeedbacklogic.UpdateAdvancedBlessActivityData(pl, int32(advanceId), advancedType)
	return
}

func init() {
	gameevent.AddEventListener(xiantiventtypes.EventTypeXianTiAdvanced, event.EventListenerFunc(playerXianTiAdavancedRew))
}
