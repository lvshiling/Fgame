package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	advancedrewrewextendedlogic "fgame/fgame/game/welfare/advancedrew/rew_extended/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
	wingeventtypes "fgame/fgame/game/wing/event/types"
)

//玩家战翼进阶
func playerWingAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId, ok := data.(int)
	if !ok {
		return
	}

	advancedType := welfaretypes.AdvancedTypeWing
	advancedrewrewextendedlogic.UpdateAdvancedRewExtendedData(pl, int32(advanceId), advancedType)
	return
}

func init() {
	gameevent.AddEventListener(wingeventtypes.EventTypeWingAdvanced, event.EventListenerFunc(playerWingAdavanced))
}
