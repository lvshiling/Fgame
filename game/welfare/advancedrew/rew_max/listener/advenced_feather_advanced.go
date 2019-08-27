package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	advancedrewrewlogic "fgame/fgame/game/welfare/advancedrew/rew_max/logic"
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
	advancedrewrewlogic.UpdateAdvancedRewData(pl, int32(advanceId), advancedType)
	return
}

func init() {
	gameevent.AddEventListener(wingeventtypes.EventTypeFeatherAdvanced, event.EventListenerFunc(playerFeatherAdavanced))
}
