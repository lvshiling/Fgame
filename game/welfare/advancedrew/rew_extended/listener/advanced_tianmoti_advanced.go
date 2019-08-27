package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	tianmotiventtypes "fgame/fgame/game/tianmo/event/types"
	advancedrewrewextendedlogic "fgame/fgame/game/welfare/advancedrew/rew_extended/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家天魔体进阶
func playerTianMoTiAdavancedRew(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId, ok := data.(int32)
	if !ok {
		return
	}

	advancedType := welfaretypes.AdvancedTypeTianMoTi
	advancedrewrewextendedlogic.UpdateAdvancedRewExtendedData(pl, int32(advanceId), advancedType)
	return
}

func init() {
	gameevent.AddEventListener(tianmotiventtypes.EventTypeTianMoAdvanced, event.EventListenerFunc(playerTianMoTiAdavancedRew))
}
