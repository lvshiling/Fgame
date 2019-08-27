package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	mounteventtypes "fgame/fgame/game/mount/event/types"
	"fgame/fgame/game/player"
	advancedrewrewextendedlogic "fgame/fgame/game/welfare/advancedrew/rew_extended/logic"
	welfaretypes "fgame/fgame/game/welfare/types"
)

//玩家坐骑进阶
func playerMountAdavanced(target event.EventTarget, data event.EventData) (err error) {
	pl, ok := target.(player.Player)
	if !ok {
		return
	}
	advanceId, ok := data.(int)
	if !ok {
		return
	}
	advancedType := welfaretypes.AdvancedTypeMount
	advancedrewrewextendedlogic.UpdateAdvancedRewExtendedData(pl, int32(advanceId), advancedType)
	return
}

func init() {
	gameevent.AddEventListener(mounteventtypes.EventTypeMountAdvanced, event.EventListenerFunc(playerMountAdavanced))
}
